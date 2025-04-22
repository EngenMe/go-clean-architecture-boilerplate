package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
	"golang.org/x/crypto/bcrypt"
)

// UpdateUserCommand is a command to update an existing user
type UpdateUserCommand struct {
	ID        uint   `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password"` // Optional
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UpdateUserHandler handles updating of users
type UpdateUserHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the update user command
func (h *UpdateUserHandler) Handle(
	ctx context.Context,
	command UpdateUserCommand,
) (*entities.UserDTO, error) {
	// Check if user exists
	user, err := h.UserRepository.GetByID(ctx, command.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrNotFound
	}

	// Check if email has changed and is already taken by someone else
	if user.Email != command.Email {
		existingUser, err := h.UserRepository.GetByEmail(ctx, command.Email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil && existingUser.ID != command.ID {
			return nil, utils.ErrConflict
		}
	}

	// Update user fields
	user.Email = command.Email
	user.FirstName = command.FirstName
	user.LastName = command.LastName
	user.UpdatedAt = time.Now()

	// Update password if provided
	if command.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(command.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := h.UserRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return &userDTO, nil
}

// RegisterUpdateUserHandler registers the update user command handler
func RegisterUpdateUserHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[UpdateUserCommand, *entities.UserDTO](
		&UpdateUserHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register UpdateUserHandler: %w", err)
	}

	return nil
}
