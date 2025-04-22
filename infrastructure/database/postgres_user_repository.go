package database

import (
	"context"
	"errors"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"gorm.io/gorm"
)

// PostgresUserRepository implements UserRepository interface using PostgreSQL
type PostgresUserRepository struct {
	db *gorm.DB
}

// NewPostgresUserRepository creates a new PostgreSQL user repository
func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &PostgresUserRepository{db: db}
}

// Create adds a new user to the database
func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user *entities.User,
) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetByID retrieves a user by ID
func (r *PostgresUserRepository) GetByID(
	ctx context.Context,
	id uint,
) (*entities.User, error) {
	var user entities.User
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No user found
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *PostgresUserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*entities.User, error) {
	var user entities.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No user found
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetAll retrieves all users
func (r *PostgresUserRepository) GetAll(ctx context.Context) (
	[]entities.User,
	error,
) {
	var users []entities.User
	result := r.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Update updates an existing user
func (r *PostgresUserRepository) Update(
	ctx context.Context,
	user *entities.User,
) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete removes a user by ID
func (r *PostgresUserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.User{}, id).Error
}
