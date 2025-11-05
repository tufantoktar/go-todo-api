package main

import (
    "log"
    "net/http"
    "os"

    "github.com/yourname/go-todo-api/internal/server"
    "github.com/yourname/go-todo-api/internal/todo"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    s := todo.NewStore()
    r := server.NewRouter(s)

    log.Printf("listening on :%s", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatal(err)
    }
}
