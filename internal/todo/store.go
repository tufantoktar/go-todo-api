package todo

import (
    "errors"
    "sync"

    "github.com/google/uuid"
)

var ErrNotFound = errors.New("todo not found")

type Store struct {
    mu    sync.RWMutex
    items map[string]Todo
}

func NewStore() *Store {
    return &Store{items: make(map[string]Todo)}
}

func (s *Store) List() []Todo {
    s.mu.RLock()
    defer s.mu.RUnlock()
    out := make([]Todo, 0, len(s.items))
    for _, t := range s.items {
        out = append(out, t)
    }
    return out
}

func (s *Store) Create(title string) Todo {
    s.mu.Lock()
    defer s.mu.Unlock()
    id := uuid.NewString()
    t := Todo{ID: id, Title: title, Completed: false}
    s.items[id] = t
    return t
}

func (s *Store) Get(id string) (Todo, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    t, ok := s.items[id]
    if !ok {
        return Todo{}, ErrNotFound
    }
    return t, nil
}

func (s *Store) Update(id string, title *string, completed *bool) (Todo, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    t, ok := s.items[id]
    if !ok {
        return Todo{}, ErrNotFound
    }
    if title != nil {
        t.Title = *title
    }
    if completed != nil {
        t.Completed = *completed
    }
    s.items[id] = t
    return t, nil
}

func (s *Store) Delete(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, ok := s.items[id]; !ok {
        return ErrNotFound
    }
    delete(s.items, id)
    return nil
}
