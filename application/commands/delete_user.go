package commands

import (
	"context"
	"fmt"

	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
)

// DeleteUserCommand is a command to delete a user
type DeleteUserCommand struct {
	ID uint `json:"id" binding:"required" example:"1"`
}

// DeleteUserHandler handles deletion of users
type DeleteUserHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the delete user command
func (h *DeleteUserHandler) Handle(
	ctx context.Context,
	command DeleteUserCommand,
) (error, error) {
	// Check if user exists
	user, err := h.UserRepository.GetByID(ctx, command.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrNotFound
	}

	err = h.UserRepository.Delete(ctx, command.ID)
	return nil, err
}

// RegisterDeleteUserHandler registers the delete user command handler
func RegisterDeleteUserHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[DeleteUserCommand, error](
		&DeleteUserHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register DeleteUserHandler: %w", err)
	}

	return nil
}
