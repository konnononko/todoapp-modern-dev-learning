package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetTodos(t *testing.T) {
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
	todos = []Todo{}
	nextID = 1

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
