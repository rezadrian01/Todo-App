package service

import (
	"context"
	"errors"
	"testing"

	"github.com/industrix-todo/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo_Success(t *testing.T) {
	repo := &mockTodoRepo{}
	svc := NewTodoService(repo)

	todo := &domain.Todo{Title: "Test", Priority: "high"}
	err := svc.CreateTodo(context.Background(), todo)
	
	assert.NoError(t, err)
	assert.Len(t, repo.todos, 1)
	assert.Equal(t, "Test", repo.todos[0].Title)
	assert.Equal(t, "high", repo.todos[0].Priority)
}

func TestCreateTodo_EmptyTitle(t *testing.T) {
	repo := &mockTodoRepo{}
	svc := NewTodoService(repo)

	err := svc.CreateTodo(context.Background(), &domain.Todo{Title: ""})
	assert.ErrorContains(t, err, "title is required")
}

func TestCreateTodo_InvalidPriority(t *testing.T) {
	repo := &mockTodoRepo{}
	svc := NewTodoService(repo)

	err := svc.CreateTodo(context.Background(), &domain.Todo{Title: "Test", Priority: "urgent"})
	assert.ErrorContains(t, err, "invalid priority")
}

func TestDeleteTodo_NotFound(t *testing.T) {
	repo := &mockTodoRepo{err: errors.New("not found")}
	svc := NewTodoService(repo)

	err := svc.DeleteTodo(context.Background(), 999)
	assert.Error(t, err)
}

func TestToggleComplete(t *testing.T) {
    repo := &mockTodoRepo{
        todos: []domain.Todo{
            {ID: 1, Title: "Test", Completed: false},
        },
    }
    svc := NewTodoService(repo)
    
    err := svc.ToggleComplete(context.Background(), 1)
    assert.NoError(t, err)
    assert.True(t, repo.todos[0].Completed)
}
