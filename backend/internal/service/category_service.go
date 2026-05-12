package service

import (
	"context"
	"errors"
	"github.com/industrix-todo/backend/internal/domain"
)

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) ListCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.List(ctx)
}

func (s *categoryService) CreateCategory(ctx context.Context, category *domain.Category) error {
	if category.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id uint, category *domain.Category) error {
	if category.Name == "" {
		return errors.New("name is required")
	}
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("category not found")
	}
	existing.Name = category.Name
	existing.Color = category.Color
	return s.repo.Update(ctx, existing)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.repo.Delete(ctx, id)
}
