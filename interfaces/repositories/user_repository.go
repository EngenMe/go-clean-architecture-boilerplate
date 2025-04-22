package repositories

import (
	"context"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
)

// UserRepository defines operations for user storage
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
}
