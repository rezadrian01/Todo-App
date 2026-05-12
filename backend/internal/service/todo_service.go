package service

import (
	"context"
	"errors"
	"github.com/industrix-todo/backend/internal/domain"
)

type todoService struct {
	repo domain.TodoRepository
}

func NewTodoService(repo domain.TodoRepository) domain.TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) ListTodos(ctx context.Context, filter domain.TodoFilter) ([]domain.Todo, int64, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	return s.repo.List(ctx, filter)
}

func (s *todoService) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	if todo.Priority == "" {
		todo.Priority = "medium"
	}
	if todo.Priority != "high" && todo.Priority != "medium" && todo.Priority != "low" {
		return errors.New("invalid priority")
	}
	return s.repo.Create(ctx, todo)
}

func (s *todoService) GetTodo(ctx context.Context, id uint) (*domain.Todo, error) {
	todo, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id uint, todo *domain.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	if todo.Priority != "high" && todo.Priority != "medium" && todo.Priority != "low" {
		return errors.New("invalid priority")
	}
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("todo not found")
	}
	existing.Title = todo.Title
	existing.Description = todo.Description
	existing.Completed = todo.Completed
	existing.Priority = todo.Priority
	existing.DueDate = todo.DueDate
	existing.CategoryID = todo.CategoryID
	return s.repo.Update(ctx, existing)
}

func (s *todoService) DeleteTodo(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("todo not found")
	}
	return s.repo.Delete(ctx, id)
}

func (s *todoService) ToggleComplete(ctx context.Context, id uint) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("todo not found")
	}
	existing.Completed = !existing.Completed
	return s.repo.Update(ctx, existing)
}
