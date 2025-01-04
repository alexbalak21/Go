package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Todo represents a single todo item
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var (
	todos  []Todo     // slice to store todo items
	nextID int        // variable to keep track of the next ID to assign
	mu     sync.Mutex // mutex to ensure thread-safe access to the todos slice
)

// init initializes the todos slice and nextID
func init() {
	todos = []Todo{}
	nextID = 1
}

// AddTodoHandler handles adding a new todo item
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a Todo struct
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lock the mutex to ensure thread-safe access to the todos slice
	mu.Lock()
	// Assign an ID to the new todo and increment nextID
	todo.ID = nextID
	nextID++
	// Append the new todo to the todos slice
	todos = append(todos, todo)
	// Unlock the mutex
	mu.Unlock()

	// Respond with the created status and the new todo item
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// ListTodosHandler handles listing all todo items
func ListTodosHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Lock the mutex to ensure thread-safe access to the todos slice
	mu.Lock()
	defer mu.Unlock()
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode the todos slice to JSON and write it to the response
	json.NewEncoder(w).Encode(todos)
}

// DeleteTodoHandler handles deleting a todo item
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a Todo struct
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lock the mutex to ensure thread-safe access to the todos slice
	mu.Lock()
	defer mu.Unlock()
	// Find the todo item with the given ID and remove it from the slice
	for i, t := range todos {
		if t.ID == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	// Respond with no content status
	w.WriteHeader(http.StatusNoContent)
}
