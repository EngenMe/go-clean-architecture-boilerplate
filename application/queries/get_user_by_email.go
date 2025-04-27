package queries

import (
	"context"
	"fmt"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
)

// GetUserByEmailQuery is a query to get a user by email
type GetUserByEmailQuery struct {
	Email string `json:"email" binding:"required,email" example:"user@example.com"`
}

// GetUserByEmailHandler handles retrieving a user by email
type GetUserByEmailHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the get user by email query
func (h *GetUserByEmailHandler) Handle(
	ctx context.Context,
	query GetUserByEmailQuery,
) (*entities.UserDTO, error) {
	user, err := h.UserRepository.GetByEmail(ctx, query.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrNotFound
	}

	userDTO := user.ToDTO()
	return &userDTO, nil
}

// RegisterGetUserByEmailHandler registers the get user by email query handler
func RegisterGetUserByEmailHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[GetUserByEmailQuery, *entities.UserDTO](
		&GetUserByEmailHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register GetUserByEmailHandler: %w", err)
	}

	return nil
}
