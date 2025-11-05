package todo

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
)

type Handlers struct{ store *Store }

func NewHandlers(s *Store) *Handlers { return &Handlers{store: s} }

func (h *Handlers) List(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, h.store.List())
}

type createReq struct {
    Title string `json:"title"`
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
    var req createReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    t := h.store.Create(req.Title)
    writeJSON(w, http.StatusCreated, t)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    t, err := h.store.Get(id)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    writeJSON(w, http.StatusOK, t)
}

type updateReq struct {
    Title     *string `json:"title,omitempty"`
    Completed *bool   `json:"completed,omitempty"`
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var req updateReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    t, err := h.store.Update(id, req.Title, req.Completed)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    writeJSON(w, http.StatusOK, t)
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    if err := h.store.Delete(id); err != nil {
        http.NotFound(w, r)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    _ = json.NewEncoder(w).Encode(v)
}
