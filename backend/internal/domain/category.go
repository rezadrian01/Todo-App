package domain

import (
	"context"
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;unique"`
	Color     string    `json:"color" gorm:"default:'#3B82F6'"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryRepository interface {
	List(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category *Category) error
	FindByID(ctx context.Context, id uint) (*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uint) error
}

type CategoryService interface {
	ListCategories(ctx context.Context) ([]Category, error)
	CreateCategory(ctx context.Context, category *Category) error
	UpdateCategory(ctx context.Context, id uint, category *Category) error
	DeleteCategory(ctx context.Context, id uint) error
}
