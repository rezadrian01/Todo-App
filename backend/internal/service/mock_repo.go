package service

import (
	"context"
	"errors"

	"github.com/industrix-todo/backend/internal/domain"
)

// MockCategoryRepo
type mockCategoryRepo struct {
	categories []domain.Category
	err        error
}

func (m *mockCategoryRepo) List(ctx context.Context) ([]domain.Category, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.categories, nil
}

func (m *mockCategoryRepo) Create(ctx context.Context, category *domain.Category) error {
	if m.err != nil {
		return m.err
	}
	category.ID = uint(len(m.categories) + 1)
	m.categories = append(m.categories, *category)
	return nil
}

func (m *mockCategoryRepo) FindByID(ctx context.Context, id uint) (*domain.Category, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, c := range m.categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockCategoryRepo) Update(ctx context.Context, category *domain.Category) error {
	if m.err != nil {
		return m.err
	}
	for i, c := range m.categories {
		if c.ID == category.ID {
			m.categories[i] = *category
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockCategoryRepo) Delete(ctx context.Context, id uint) error {
	if m.err != nil {
		return m.err
	}
	for i, c := range m.categories {
		if c.ID == id {
			m.categories = append(m.categories[:i], m.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

// MockTodoRepo
type mockTodoRepo struct {
	todos []domain.Todo
	err   error
}

func (m *mockTodoRepo) List(ctx context.Context, filter domain.TodoFilter) ([]domain.Todo, int64, error) {
	if m.err != nil {
		return nil, 0, m.err
	}
	return m.todos, int64(len(m.todos)), nil
}

func (m *mockTodoRepo) Create(ctx context.Context, todo *domain.Todo) error {
	if m.err != nil {
		return m.err
	}
	todo.ID = uint(len(m.todos) + 1)
	m.todos = append(m.todos, *todo)
	return nil
}

func (m *mockTodoRepo) FindByID(ctx context.Context, id uint) (*domain.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, t := range m.todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockTodoRepo) Update(ctx context.Context, todo *domain.Todo) error {
	if m.err != nil {
		return m.err
	}
	for i, t := range m.todos {
		if t.ID == todo.ID {
			m.todos[i] = *todo
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockTodoRepo) Delete(ctx context.Context, id uint) error {
	if m.err != nil {
		return m.err
	}
	for i, t := range m.todos {
		if t.ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
