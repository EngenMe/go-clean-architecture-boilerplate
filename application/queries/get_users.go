package queries

import (
	"context"
	"fmt"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
)

// GetUsersQuery is a query to get all users
type GetUsersQuery struct{}

// GetUsersHandler handles retrieving all users
type GetUsersHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the get all users query
func (h *GetUsersHandler) Handle(
	ctx context.Context,
	_ GetUsersQuery,
) ([]entities.UserDTO, error) {
	users, err := h.UserRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	userDTOs := make([]entities.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = user.ToDTO()
	}

	return userDTOs, nil
}

// RegisterGetUsersHandler registers the get all users query handler
func RegisterGetUsersHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[GetUsersQuery, []entities.UserDTO](
		&GetUsersHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register GetUserHandler: %w", err)
	}

	return nil
}
