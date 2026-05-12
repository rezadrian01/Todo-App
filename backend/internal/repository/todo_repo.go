package repository

import (
	"context"
	"github.com/industrix-todo/backend/internal/domain"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepo{db: db}
}

func (r *todoRepo) List(ctx context.Context, f domain.TodoFilter) ([]domain.Todo, int64, error) {
	var todos []domain.Todo
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Todo{})

	if f.Search != "" {
		query = query.Where("to_tsvector('english', title) @@ plainto_tsquery(?)", f.Search)
	}
	if f.Status == "completed" {
		query = query.Where("completed = ?", true)
	} else if f.Status == "incomplete" {
		query = query.Where("completed = ?", false)
	}
	if f.CategoryID != 0 {
		query = query.Where("category_id = ?", f.CategoryID)
	}
	if f.Priority != "" {
		query = query.Where("priority = ?", f.Priority)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Preload("Category")
    
	sortBy := "created_at"
	if f.SortBy != "" {
		sortBy = f.SortBy
	}
	sortOrder := "DESC"
	if f.SortOrder != "" {
		sortOrder = f.SortOrder
	}
	
	orderQuery := sortBy + " " + sortOrder

	offset := (f.Page - 1) * f.Limit
	err := query.
		Order(orderQuery).
		Offset(offset).
		Limit(f.Limit).
		Find(&todos).Error

	return todos, total, err
}

func (r *todoRepo) Create(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepo) FindByID(ctx context.Context, id uint) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.WithContext(ctx).Preload("Category").First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepo) Update(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Todo{}, id).Error
}
