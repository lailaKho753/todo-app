package main

import (
	"fmt"
	"sync"
)

type InMemoryStore struct {
	mu     sync.Mutex
	todos  map[int]*Todo
	nextID int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		todos:  make(map[int]*Todo),
		nextID: 1,
	}
}

func (s *InMemoryStore) Create(todo *Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	s.nextID++

	s.todos[todo.ID] = todo
	return nil
}

func (s *InMemoryStore) GetAll() ([]Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Todo
	for _, t := range s.todos {
		result = append(result, *t)
	}

	return result, nil
}

func (s *InMemoryStore) GetByID(id int) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	return todo, nil
}

func (s *InMemoryStore) Update(id int, updated *Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return fmt.Errorf("todo not found")
	}

	updated.ID = id
	s.todos[id] = updated
	return nil
}

func (s *InMemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return fmt.Errorf("todo not found")
	}

	delete(s.todos, id)
	return nil
}

