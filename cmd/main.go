package main

import (
	"go-api/internal/handlers"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", handlers.ListTodosHandler)
	mux.HandleFunc("/todos/add", handlers.AddTodoHandler)
	mux.HandleFunc("/todos/delete", handlers.DeleteTodoHandler)

	loggedMux := loggingMiddleware(mux)

	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
