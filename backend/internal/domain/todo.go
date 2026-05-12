package domain

import (
	"context"
	"time"
)

type Todo struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed" gorm:"default:false"`
	Priority    string     `json:"priority" gorm:"default:'medium'"`
	DueDate     *time.Time `json:"due_date"`
	CategoryID  *uint      `json:"category_id"`
	Category    *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TodoFilter struct {
	Page       int
	Limit      int
	Search     string
	Status     string // "completed" | "incomplete"
	CategoryID uint
	Priority   string
	SortBy     string
	SortOrder  string
}

type TodoRepository interface {
	List(ctx context.Context, filter TodoFilter) ([]Todo, int64, error)
	Create(ctx context.Context, todo *Todo) error
	FindByID(ctx context.Context, id uint) (*Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id uint) error
}

type TodoService interface {
	ListTodos(ctx context.Context, filter TodoFilter) ([]Todo, int64, error)
	CreateTodo(ctx context.Context, todo *Todo) error
	GetTodo(ctx context.Context, id uint) (*Todo, error)
	UpdateTodo(ctx context.Context, id uint, todo *Todo) error
	DeleteTodo(ctx context.Context, id uint) error
	ToggleComplete(ctx context.Context, id uint) error
}
