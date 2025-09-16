# Go HTMX Templ Blueprint - Makefile
.PHONY: help install dev build clean test templ css watch-css deps air-install tailwind-install

# Default target
help: ## Show this help message
	@echo "Go HTMX Templ Blueprint"
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Installation targets
install: deps tailwind-install ## Install all dependencies
	@echo "✅ All dependencies installed successfully!"

deps: ## Install Go dependencies
	@echo "📦 Installing Go dependencies..."
	go mod tidy
	go mod download


tailwind-install: ## Install Tailwind CSS CLI
	@echo "🎨 Installing Tailwind CSS CLI..."
	@if ! command -v tailwindcss > /dev/null 2>&1; then \
		echo "Detecting OS..."; \
		OS=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
		ARCH=$$(uname -m); \
		if [ "$$ARCH" = "x86_64" ]; then ARCH="x64"; fi; \
		if [ "$$ARCH" = "aarch64" ] || [ "$$ARCH" = "arm64" ]; then ARCH="arm64"; fi; \
		URL="https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$$OS-$$ARCH"; \
		echo "Downloading from: $$URL"; \
		curl -sLO $$URL; \
		chmod +x tailwindcss-$$OS-$$ARCH; \
		sudo mv tailwindcss-$$OS-$$ARCH /usr/local/bin/tailwindcss; \
		echo "✅ Tailwind CSS CLI installed"; \
	else \
		echo "✅ Tailwind CSS CLI already installed"; \
	fi

# Development targets
dev: ## Start development server with hot reloading
	@echo "🚀 Starting development server..."
	@make -j2 watch-css templ-watch

templ-watch: ## Watch and generate templ files
	@echo "👀 Watching templ files..."
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ./cmd/server/" --open-browser=false -v

watch-css: ## Watch and build CSS
	@echo "🎨 Watching CSS files..."
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch -v

# Build targets
build: templ css ## Build the application
	@echo "🔨 Building application..."
	go build -o bin/app ./cmd/server

templ: ## Generate templ files
	@echo "📝 Generating templ files..."
	templ generate

css: ## Build CSS
	@echo "🎨 Building CSS..."
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

# Production build
build-prod: templ css ## Build for production
	@echo "🏗️  Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app ./cmd/server

# Utility targets
clean: ## Clean build artifacts
	@echo "🧹 Cleaning up..."
	rm -rf bin/
	rm -f static/css/output.css

test: ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

fmt: ## Format Go code
	@echo "📐 Formatting code..."
	go fmt ./...
	templ fmt .

lint: ## Run linters
	@echo "🔍 Running linters..."
	golangci-lint run

setup-dirs: ## Create necessary directories
	@echo "📁 Creating directories..."
	mkdir -p bin tmp static/css static/js templates components handlers

# Development helpers
logs: ## Show air logs
	tail -f tmp/air.log

restart: ## Restart development server
	@echo "🔄 Restarting development server..."
	@pkill -f "air" || true
	@pkill -f "tailwindcss.*watch" || true
	@pkill -f "templ.*watch" || true
	@make dev
