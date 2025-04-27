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
	Email     string `json:"email" binding:"required,email" example:"user@example.com"`
	Password  string `json:"password" binding:"required,min=6" example:"password123"`
	FirstName string `json:"firstName" binding:"required" example:"John"`
	LastName  string `json:"lastName" binding:"required" example:"Doe"`
}

// CreateUserHandler handle creation of new users
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
		return nil, utils.ErrEmailAlreadyExists
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(command.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
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
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	userDTO := user.ToDTO()
	return &userDTO, nil
}

// RegisterCreateUserHandler registers the creation user command handler
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
