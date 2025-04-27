package queries

import (
	"context"
	"fmt"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
)

// GetUserByIDQuery is a query to get a user by ID
type GetUserByIDQuery struct {
	ID uint `json:"id" binding:"required" example:"1"`
}

// GetUserByIDHandler handles retrieving a user by ID
type GetUserByIDHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the get user by ID query
func (h *GetUserByIDHandler) Handle(
	ctx context.Context,
	query GetUserByIDQuery,
) (*entities.UserDTO, error) {
	user, err := h.UserRepository.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrNotFound
	}

	userDTO := user.ToDTO()
	return &userDTO, nil
}

// RegisterGetUserByIDHandler registers the get user by ID query handler
func RegisterGetUserByIDHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[GetUserByIDQuery, *entities.UserDTO](
		&GetUserByIDHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register GetUserByIDHandler: %w", err)
	}

	return nil
}
