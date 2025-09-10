# Go HTMX Templ Blueprint

A modern, server-driven web application blueprint using Go, HTMX, Templ, and Tailwind CSS. This project demonstrates the HATEOAS pattern for building interactive web applications with minimal JavaScript.

## 🚀 Tech Stack

- **Backend**: Go with Chi router
- **Frontend**: HTMX for interactivity
- **Templating**: Templ for type-safe templates
- **Styling**: Tailwind CSS v4
- **Hot Reload**: Air for Go, built-in watchers for CSS and templates

## 📁 Project Structure

```
your-project/
├── cmd/server/           # Application entrypoint
│   └── main.go
├── internal/             # Private application code
│   ├── config/           # Configuration management
│   ├── handlers/         # HTTP handlers
│   ├── components/       # Templ components
│   └── server/           # Server setup and routing
├── templates/            # Page templates
├── static/               # Static assets
│   ├── css/
│   └── js/
├── bin/                  # Compiled binaries
├── tmp/                  # Temporary files (Air)
├── Makefile              # Build automation
├── .air.toml            # Air configuration
├── tailwind.config.js   # Tailwind configuration
└── go.mod               # Go module definition
```

## 🛠️ Installation

### Prerequisites

- Go 1.21 or later
- Make (optional, but recommended)

### Quick Start

1. **Clone and setup the project:**
   ```bash
   git clone <your-repo>
   cd your-project
   ```

2. **Install all dependencies:**
   ```bash
   make install
   ```
   This will install:
   - Go dependencies
   - Air (hot reloading)
   - Tailwind CSS CLI

3. **Start development server:**
   ```bash
   make dev
   ```
   This starts three processes concurrently:
   - Air (Go hot reloading)
   - Templ watcher (template generation)
   - Tailwind CSS watcher (CSS compilation)

4. **Visit your application:**
   Open [http://localhost:8080](http://localhost:8080)

### Manual Installation

If you prefer to install dependencies manually:

```bash
# Install Go dependencies
go mod tidy

# Install Air
go install github.com/cosmtrek/air@latest

# Install Tailwind CSS CLI (Linux/macOS)
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
sudo mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Install Templ
go install github.com/a-h/templ/cmd/templ@latest
```

## 🧰 Available Commands

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make install` | Install all dependencies |
| `make dev` | Start development server with hot reload |
| `make build` | Build the application |
| `make build-prod` | Build for production |
| `make test` | Run tests |
| `make clean` | Clean build artifacts |
| `make fmt` | Format code |

## 🎯 Features Demonstrated

### HTMX Patterns
- **hx-get/post/put/delete**: RESTful actions without page refresh
- **hx-target**: Specify where to update content
- **hx-indicator**: Loading states and feedback
- **hx-confirm**: User confirmation dialogs
- **hx-vals**: Dynamic form values

### Server-Side Rendering
- Type-safe templates with Templ
- Component-based architecture
- Partial page updates
- HATEOAS-compliant responses

### Modern Development Experience
- Hot reloading for Go, templates, and CSS
- Tailwind CSS with custom components
- Structured logging
- Graceful shutdown
- Health check endpoints

## 🏗️ Architecture Patterns

### HATEOAS (Hypermedia as the Engine of Application State)
This application follows HATEOAS principles:
- Server responses include the hypermedia controls (HTML)
- Client state is driven by server responses
- Navigation and actions are embedded in the response
- Minimal client-side state management

### Component Architecture
```
Page Templates (templates/)
    ↓
Base Layout (templates/base.templ)
    ↓
Reusable Components (internal/components/)
    ↓
Handlers (internal/handlers/)
```

### Request Flow
```
HTTP Request → Chi Router → Handler → Templ Component → HTML Response
```

## 🔧 Configuration

### Environment Variables
Copy `.env.example` to `.env` and customize:

```bash
ENVIRONMENT=development
PORT=8080
DATABASE_URL=your_database_url
```

### Tailwind Configuration
Customize `tailwind.config.js` to match your design system:
- Add custom colors
- Configure fonts
- Add plugins
- Set up custom components

### Air Configuration
Modify `.air.toml` to change:
- File watching patterns
- Build commands
- Excluded directories

## 🚀 Deployment

### Build for Production
```bash
make build-prod
```

### Docker
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN make build-prod

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/app .
COPY --from=builder /app/static ./static
CMD ["./app"]
```

## 📚 Learning Resources

### HTMX
- [HTMX Documentation](https://htmx.org/docs/)
- [HTMX Examples](https://htmx.org/examples/)

### Templ
- [Templ Documentation](https://templ.guide/)
- [Templ GitHub](https://github.com/a-h/templ)

### Go Chi
- [Chi Documentation](https://go-chi.io/)
- [Chi GitHub](https://github.com/go-chi/chi)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Format code: `make fmt`
6. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🎉 What's Next?

This blueprint provides a solid foundation. Consider adding:
- Database integration (PostgreSQL, SQLite)
- Authentication and authorization
- WebSocket support for real-time features
- API versioning
- Metrics and monitoring
- Docker compose for local development
- CI/CD pipeline configuration

Happy coding with HTMX and Go! 🚀
