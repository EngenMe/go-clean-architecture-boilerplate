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

// CreateUserCommand is a command to create a new user
type CreateUserCommand struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// CreateUserHandler handles creation of new users
type CreateUserHandler struct {
	UserRepository repositories.UserRepository
}

// Handle processes the create user command
func (h *CreateUserHandler) Handle(
	ctx context.Context,
	command CreateUserCommand,
) (*entities.UserDTO, error) {
	// Check if user already exists
	existingUser, err := h.UserRepository.GetByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, utils.ErrConflict
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(command.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &entities.User{
		Email:     command.Email,
		Password:  string(hashedPassword),
		FirstName: command.FirstName,
		LastName:  command.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.UserRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return &userDTO, nil
}

// RegisterCreateUserHandler registers the create user command handler
func RegisterCreateUserHandler(userRepository repositories.UserRepository) error {
	if err := mediatr.RegisterRequestHandler[CreateUserCommand, *entities.UserDTO](
		&CreateUserHandler{
			UserRepository: userRepository,
		},
	); err != nil {
		return fmt.Errorf("failed to register CreateUserHandler: %w", err)
	}

	return nil
}
