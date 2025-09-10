package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	logger *slog.Logger
	todos  []Todo // in-memory store for blueprint demo pourpose
}

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func New(logger *slog.Logger) *Handlers {
	return &Handlers{
		logger: logger,
		todos: []Todo{
			{ID: 1, Title: "Learn HTMX", Completed: false, CreatedAt: time.Now()},
			{ID: 2, Title: "Build with Templ", Completed: false, CreatedAt: time.Now()},
			{ID: 3, Title: "Beautyfy with tailwindcss", Completed: false, CreatedAt: time.Now()},
			{ID: 4, Title: "Ship with Go", Completed: false, CreatedAt: time.Now()},
		},
	}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "ok", "timestamp":`+time.Now().Format(time.RFC3339)+`"}`)
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	page := templates.HomePage(h.todos)
	err := page.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering Home page", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) GetTodos(w http.ResponseWriter, r *http.Request) {
	component := components.TodoList(h.todos)
	err := component.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering the todos", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	if title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	newTodo := Todo{
		ID:        len(h.todos) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	h.todos = append(h.todos, newTodo)

	comp := components.TodoList(h.todos)
	err := comp.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering the todos after the creation", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id URL Param is missing or invalid", http.StatusBadRequest)
		return
	}

	for i, todo := range h.todos {
		if todo.ID == id {
			if title := r.FormValue("title"); title != "" {
				h.todos[i].Title = title
			}
			if completed := r.FormValue("completed"); completed != "" {
				h.todos[i].Completed = completed == "true"
			}
			break
		}
	}

	comp := components.TodoList(h.todos)
	err = comp.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering the todos after update", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id URL Param is missing or invalid", http.StatusBadRequest)
		return
	}

	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			break
		}
	}

	comp := components.TodoList(h.todos)
	err = comp.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering the todos after deletion", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) EditTodoForm(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id URL Param is missing or invalid", http.StatusBadRequest)
		return
	}

	var todo *Todo
	for _, t := range h.todos {
		if t.ID == id {
			todo = &t
			break
		}
	}

	if todo == nil {
		http.Error(w, "Item Not Found", http.StatusNotFound)
		return
	}

	comp := components.EditTodoForm(*todo)
	err := comp.Render(r.Context(), w)
	if err != nil {
		h.logger.Error("Error rendering the edit form", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
