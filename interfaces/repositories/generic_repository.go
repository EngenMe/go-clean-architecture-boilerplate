package repositories

import (
	"context"

	"gorm.io/gorm"
)

// Entity defines common operations for entities
type Entity interface {
	GetID() uint
	SetID(id uint)
}

// GenericRepository defines generic CRUD operations
type GenericRepository[T Entity] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id uint) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
	GetDB() *gorm.DB
}
