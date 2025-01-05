package main

import (
	"go-api/internal/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetResponseHandler).Methods("GET")
	r.HandleFunc("/todos", handlers.PostResponseHandler).Methods("POST")
	r.HandleFunc("/todos", handlers.DeleteResponseHandler).Methods("DELETE")

	loggedRouter := loggingMiddleware(r)

	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", loggedRouter)
}
