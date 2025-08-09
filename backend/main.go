package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type PatchTodo struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/todos", getTodos)
	r.Post("/todos", createTodo)
	r.Patch("/todos/{id}", updateTodo)
	r.Delete("/todos/{id}", deleteTodo)

	http.ListenAndServe(":8080", r)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func addTodo(todo Todo) Todo {
	mu.Lock()
	defer mu.Unlock()

	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	return todo
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	created := addTodo(newTodo)
	if err := writeJSON(w, http.StatusCreated, created); err != nil {
		log.Println("createTodo writeJSON failed")
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, todo := range todos {
		if todo.ID == id {
			index = i
			break
		}
	}

	mu.Lock()
	defer mu.Unlock()

	if index == -1 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var patch PatchTodo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&patch); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if patch.Title == nil && patch.Done == nil {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	var found *Todo = nil
	for i := range todos {
		if todos[i].ID == id {
			found = &todos[i]
			break
		}
	}
	if found == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	if patch.Title != nil {
		found.Title = *patch.Title
	}
	if patch.Done != nil {
		found.Done = *patch.Done
	}
	if err := writeJSON(w, http.StatusOK, found); err != nil {
		log.Println("updateTodo writeJSON failed")
	}
}
