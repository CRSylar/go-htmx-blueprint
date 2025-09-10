package server

import (
	"github.com/CRSylar/go-htmx-blueprint/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {

	// Health check
	r.Get("/health", h.HealthCheck)

	// Home page
	r.Get("/", h.Home)

	// API for HTMX showcase
	r.Route("/api", func(route chi.Router) {
		route.Get("/todos", h.GetTodos)
		route.Post("/todos", h.CreateTodo)
		route.Put("/todos/{id}", h.UpdateTodo)
		route.Delete("/todos/{id}", h.DeleteTodo)

		// partial render example for htmx
		r.Get("/todos/{id}/edit", h.EditTodoForm)
	})

}
