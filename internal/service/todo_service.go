package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TodoService struct {
	Repo repository.TodoRepository
}

func (s *TodoService) Create(todo *model.Todo) error {
	return s.Repo.Create(todo)
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	return s.Repo.GetAll()
}

func (s *TodoService) GetByID(id string) (*model.Todo, error) {
	return s.Repo.GetByID(id)
}

func (s *TodoService) Update(todo *model.Todo) error {
	return s.Repo.Update(todo)
}

func (s *TodoService) Delete(id string) error {
	return s.Repo.Delete(id)
}