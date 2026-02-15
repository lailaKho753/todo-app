package main

type TodoStore interface {
	Create(todo *Todo) error
	GetAll() ([]Todo, error)
	GetByID(id int) (*Todo, error)
	Update(id int, todo *Todo) error
	Delete(id int) error
}

