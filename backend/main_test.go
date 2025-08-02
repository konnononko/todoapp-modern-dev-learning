package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func resetState() {
	todos = []Todo{}
	nextID = 1
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/todos", getTodos)
	r.Post("/todos", createTodo)
	r.Delete("/todos/{id}", deleteTodo)
	return r
}

func TestGetTodos(t *testing.T) {
	todos = []Todo{{ID: -1, Title: "hoge", Done: false}}
	nextID = 1

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/todos", nil)

	getTodos(rr, req)

	assertEqual(t, http.StatusOK, rr.Code)

	var got []Todo
	err := json.NewDecoder(rr.Body).Decode(&got)
	assertEqual(t, nil, err)
	assertEqual(t, len(todos), len(got))
}

func TestCreateTodo(t *testing.T) {
	resetState()

	payload := `{"title": "新しいタスク", "done": false}`
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	createTodo(rr, req)

	assertEqual(t, http.StatusCreated, rr.Code)
	var created Todo
	err := json.NewDecoder(rr.Body).Decode(&created)
	assertEqual(t, nil, err)
	assertEqual(t, 1, created.ID)
	assertEqual(t, "新しいタスク", created.Title)
	assertEqual(t, false, created.Done)
	assertEqual(t, 1, len(todos))
}

func TestCreateTodo_MultipleAdds(t *testing.T) {
	resetState()

	titles := []string{"task1", "task2", "task3"}

	for i, title := range titles {
		payload := fmt.Sprintf(`{"title":"%s","done":false}`, title)
		req := httptest.NewRequest("POST", "/todos", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		createTodo(rr, req)

		assertEqual(t, http.StatusCreated, rr.Code)
		var got Todo
		err := json.NewDecoder(rr.Body).Decode(&got)
		assertEqual(t, nil, err)
		assertEqual(t, i+1, got.ID)
		assertEqual(t, title, got.Title)
	}
	assertEqual(t, len(titles), len(todos))
}

func TestCreateTodo_BadRequest(t *testing.T) {
	todos = []Todo{{ID: -1, Title: "hoge", Done: false}}
	nextID = 1

	payload := `{bad json}`
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	createTodo(rr, req)

	assertEqual(t, http.StatusBadRequest, rr.Code)
	assertEqual(t, 1, len(todos))
}

func TestDeleteTodo(t *testing.T) {
	resetState()
	added := addTodo(Todo{Title: "task", Done: false})

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/todos/%d", added.ID), nil)
	rr := httptest.NewRecorder()

	r := setupRouter()
	r.ServeHTTP(rr, req)

	assertEqual(t, http.StatusNoContent, rr.Code)
	assertEqual(t, 0, len(todos))
}

func TestDeleteTodo_MultipleTodos(t *testing.T) {
	resetState()
	addTodo(Todo{Title: "task1", Done: false})
	todo2 := addTodo(Todo{Title: "task2", Done: false})
	addTodo(Todo{Title: "task3", Done: false})

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/todos/%d", todo2.ID), nil)
	rr := httptest.NewRecorder()

	r := setupRouter()
	r.ServeHTTP(rr, req)

	assertEqual(t, http.StatusNoContent, rr.Code)
	assertEqual(t, 2, len(todos))
	for _, todo := range todos {
		if todo.ID == todo2.ID {
			t.Errorf("id %d is not deleted", todo2.ID)
		}
	}
}

func TestDeleteTodo_InvalidID(t *testing.T) {
	resetState()
	addTodo(Todo{Title: "task1", Done: false})

	req := httptest.NewRequest("DELETE", "/todos/abc", nil)
	rr := httptest.NewRecorder()

	r := setupRouter()
	r.ServeHTTP(rr, req)

	assertEqual(t, http.StatusBadRequest, rr.Code)
	assertEqual(t, 1, len(todos))
}

func TestDeleteTodo_NotFound(t *testing.T) {
	resetState()
	addTodo(Todo{Title: "task1", Done: false})

	req := httptest.NewRequest("DELETE", "/todos/-1", nil)
	rr := httptest.NewRecorder()

	r := setupRouter()
	r.ServeHTTP(rr, req)

	assertEqual(t, http.StatusNotFound, rr.Code)
	assertEqual(t, 1, len(todos))
}
