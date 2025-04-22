package services

import (
	"context"
	"log"

	"github.com/EngenMe/go-clean-architecture/application/commands"
	"github.com/EngenMe/go-clean-architecture/application/queries"
	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
)

// UserService provides user-related functionality
type UserService struct {
	userRepository repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(
	ctx context.Context,
	command commands.CreateUserCommand,
) (*entities.UserDTO, error) {
	return mediatr.Send[commands.CreateUserCommand, *entities.UserDTO](
		ctx,
		command,
	)
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(
	ctx context.Context,
	id uint,
) (*entities.UserDTO, error) {
	return mediatr.Send[queries.GetUserByIDQuery, *entities.UserDTO](
		ctx,
		queries.GetUserByIDQuery{ID: id},
	)
}

// GetUserByEmail gets a user by email
func (s *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (*entities.UserDTO, error) {
	return mediatr.Send[queries.GetUserByEmailQuery, *entities.UserDTO](
		ctx,
		queries.GetUserByEmailQuery{Email: email},
	)
}

// GetAllUsers gets all users
func (s *UserService) GetAllUsers(ctx context.Context) (
	[]entities.UserDTO,
	error,
) {
	return mediatr.Send[queries.GetUsersQuery, []entities.UserDTO](
		ctx,
		queries.GetUsersQuery{},
	)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(
	ctx context.Context,
	command commands.UpdateUserCommand,
) (*entities.UserDTO, error) {
	return mediatr.Send[commands.UpdateUserCommand, *entities.UserDTO](
		ctx,
		command,
	)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	// Expect (error, error) from the handler
	resp, err := mediatr.Send[commands.DeleteUserCommand, error](
		ctx,
		commands.DeleteUserCommand{ID: id},
	)
	// Since resp is the first error (typically nil), return the second error
	if err != nil {
		return err
	}
	return resp // resp is the first error, which is nil unless the handler sets it
}

// RegisterUserService registers the user service and all its handlers
func RegisterUserService(userRepository repositories.UserRepository) *UserService {
	// Register command handlers
	if err := commands.RegisterCreateUserHandler(userRepository); err != nil {
		log.Fatalf("Failed to register CreateUserHandler: %v", err)
	}
	if err := commands.RegisterUpdateUserHandler(userRepository); err != nil {
		log.Fatalf("Failed to register UpdateUserHandler: %v", err)
	}
	if err := commands.RegisterDeleteUserHandler(userRepository); err != nil {
		log.Fatalf("Failed to register DeleteUserHandler: %v", err)
	}

	// Register query handlers
	if err := queries.RegisterGetUserByIDHandler(userRepository); err != nil {
		log.Fatalf("Failed to register GetUserByIDHandler: %v", err)
	}
	if err := queries.RegisterGetUserByEmailHandler(userRepository); err != nil {
		log.Fatalf("Failed to register GetUserByEmailHandler: %v", err)
	}
	if err := queries.RegisterGetUsersHandler(userRepository); err != nil {
		log.Fatalf("Failed to register GetUsersHandler: %v", err)
	}

	return NewUserService(userRepository)
}
