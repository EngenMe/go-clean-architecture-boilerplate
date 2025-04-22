package repositories

import (
	"context"
)

// Entity represents a generic entity with an ID
type Entity interface {
	GetID() uint
	SetID(id uint)
	TableName() string
}

// GenericRepository defines a generic interface for data operations
// Note: We're using *T to specify that we work with pointers to entities
type GenericRepository[T Entity] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id uint) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
}
