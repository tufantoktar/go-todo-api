package server

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/yourname/go-todo-api/internal/todo"
)

func NewRouter(store *todo.Store) http.Handler {
    r := chi.NewRouter()

    r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })

    h := todo.NewHandlers(store)
    r.Route("/v1/todos", func(r chi.Router) {
        r.Get("/", h.List)
        r.Post("/", h.Create)
        r.Get("/{id}", h.Get)
        r.Put("/{id}", h.Update)
        r.Delete("/{id}", h.Delete)
    })

    return r
}
