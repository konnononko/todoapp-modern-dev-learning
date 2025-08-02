package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
