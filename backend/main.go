package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{}
var nextID = 1

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/todos", getTodos)
	r.Post("/todos", createTodo)
	r.Delete("/todos/{id}", deleteTodo)

	http.ListenAndServe(":8080", r)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func addTodo(todo Todo) Todo {
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

	if index == -1 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}
