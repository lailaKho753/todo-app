package main

import "testing"

func TestCreateAndGetByID(t *testing.T) {
	store := NewInMemoryStore()

	todo := &Todo{
		Title:     "belajar testing",
		Completed: false,
	}

	err := store.Create(todo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := store.GetByID(todo.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Title != todo.Title {
		t.Errorf("expected title %s but got %s", todo.Title, got.Title)
	}
}
