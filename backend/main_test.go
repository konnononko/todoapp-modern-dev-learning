package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/todos", nil)

	getTodos(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}

	var got []Todo
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Errorf("failed to decode JSON: %v", err)
	}

	if len(got) != len(todos) {
		t.Errorf("expected %d todos, got %d", len(todos), len(got))
	}
}
