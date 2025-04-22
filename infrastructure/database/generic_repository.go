package database

import (
	"context"
	"errors"

	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"gorm.io/gorm"
)

// GenericPostgresRepository implements GenericRepository interface using PostgreSQL
type GenericPostgresRepository[T repositories.Entity] struct {
	db *gorm.DB
}

// NewGenericPostgresRepository creates a new generic PostgreSQL repository
func NewGenericPostgresRepository[T repositories.Entity](db *gorm.DB) repositories.GenericRepository[T] {
	return &GenericPostgresRepository[T]{db: db}
}

// Create adds a new entity to the database
func (r *GenericPostgresRepository[T]) Create(
	ctx context.Context,
	entity *T,
) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// FindByID retrieves an entity by ID
func (r *GenericPostgresRepository[T]) FindByID(
	ctx context.Context,
	id uint,
) (*T, error) {
	var entity T
	result := r.db.WithContext(ctx).First(&entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No entity found
		}
		return nil, result.Error
	}
	return &entity, nil
}

// FindAll retrieves all entities
func (r *GenericPostgresRepository[T]) FindAll(ctx context.Context) (
	[]*T,
	error,
) {
	var entities []*T
	result := r.db.WithContext(ctx).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

// Update updates an existing entity
func (r *GenericPostgresRepository[T]) Update(
	ctx context.Context,
	entity *T,
) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

// Delete removes an entity by ID
func (r *GenericPostgresRepository[T]) Delete(
	ctx context.Context,
	id uint,
) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}

// GetDB returns the database connection
func (r *GenericPostgresRepository[T]) GetDB() *gorm.DB {
	return r.db
}
