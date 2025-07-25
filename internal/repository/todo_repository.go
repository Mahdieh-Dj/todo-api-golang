package repository

import "todo-app/internal/model"

type TodoRepository interface {
	Create(todo *model.Todo) error
	GetAll()([]model.Todo,error)
	GetByID(id string)(*model.Todo,error)
	Update(todo *model.Todo)error
	Delete(id string)error
}