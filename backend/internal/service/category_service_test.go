package service

import (
	"context"
	"errors"
	"testing"

	"github.com/industrix-todo/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory_Success(t *testing.T) {
	repo := &mockCategoryRepo{}
	svc := NewCategoryService(repo)

	err := svc.CreateCategory(context.Background(), &domain.Category{Name: "Work"})
	assert.NoError(t, err)
	assert.Len(t, repo.categories, 1)
	assert.Equal(t, "Work", repo.categories[0].Name)
}

func TestCreateCategory_EmptyName(t *testing.T) {
	repo := &mockCategoryRepo{}
	svc := NewCategoryService(repo)

	err := svc.CreateCategory(context.Background(), &domain.Category{Name: ""})
	assert.ErrorContains(t, err, "name is required")
	assert.Len(t, repo.categories, 0)
}

func TestDeleteCategory_NotFound(t *testing.T) {
	repo := &mockCategoryRepo{err: errors.New("not found")}
	svc := NewCategoryService(repo)

	err := svc.DeleteCategory(context.Background(), 999)
	assert.Error(t, err)
}
