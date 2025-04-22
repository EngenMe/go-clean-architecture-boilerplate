package services

import (
	"context"
	"errors"
	"log"

	"github.com/EngenMe/go-clean-architecture/application/commands"
	"github.com/EngenMe/go-clean-architecture/application/queries"
	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/database"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

// UserService provides user-related functionality
type UserService struct {
	userRepository repositories.GenericRepository[entities.User]
}

// NewUserService creates a new user service
func NewUserService(userRepository repositories.GenericRepository[entities.User]) *UserService {
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
func RegisterUserService(userRepository repositories.GenericRepository[entities.User]) *UserService {
	// Create custom repository adapter if needed for existing handlers,
	// This adapter allows existing handlers to use the GenericRepository
	userRepositoryAdapter := NewUserRepositoryAdapter(userRepository)

	// Register command handlers
	if err := commands.RegisterCreateUserHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register CreateUserHandler: %v", err)
	}
	if err := commands.RegisterUpdateUserHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register UpdateUserHandler: %v", err)
	}
	if err := commands.RegisterDeleteUserHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register DeleteUserHandler: %v", err)
	}

	// Register query handlers
	if err := queries.RegisterGetUserByIDHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register GetUserByIDHandler: %v", err)
	}
	if err := queries.RegisterGetUserByEmailHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register GetUserByEmailHandler: %v", err)
	}
	if err := queries.RegisterGetUsersHandler(userRepositoryAdapter); err != nil {
		log.Fatalf("Failed to register GetUsersHandler: %v", err)
	}

	return NewUserService(userRepository)
}

// UserRepositoryAdapter adapts GenericRepository to legacy UserRepository interface
type UserRepositoryAdapter struct {
	genericRepo repositories.GenericRepository[entities.User]
}

// NewUserRepositoryAdapter creates a new adapter
func NewUserRepositoryAdapter(genericRepo repositories.GenericRepository[entities.User]) repositories.UserRepository {
	return &UserRepositoryAdapter{
		genericRepo: genericRepo,
	}
}

// Create implements UserRepository.Create
func (a *UserRepositoryAdapter) Create(
	ctx context.Context,
	user *entities.User,
) error {
	return a.genericRepo.Create(ctx, user)
}

// GetByID implements UserRepository.GetByID
func (a *UserRepositoryAdapter) GetByID(
	ctx context.Context,
	id uint,
) (*entities.User, error) {
	return a.genericRepo.FindByID(ctx, id)
}

// GetByEmail implements UserRepository.GetByEmail
func (a *UserRepositoryAdapter) GetByEmail(
	ctx context.Context,
	email string,
) (*entities.User, error) {
	// This special case needs custom implementation as it's not part of the generic repository
	// Cast to access the underlying db
	pgRepo, ok := a.genericRepo.(*database.GenericPostgresRepository[entities.User])
	if !ok {
		return nil, errors.New("failed to cast repository for GetByEmail operation")
	}

	var user entities.User
	result := pgRepo.GetDB().WithContext(ctx).Where(
		"email = ?",
		email,
	).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No user found
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetAll implements UserRepository.GetAll
func (a *UserRepositoryAdapter) GetAll(ctx context.Context) (
	[]entities.User,
	error,
) {
	// Since FindAll now returns []*T, we need to convert back for compatibility
	ptrUsers, err := a.genericRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Convert []*User to []User for backward compatibility
	users := make([]entities.User, len(ptrUsers))
	for i, u := range ptrUsers {
		if u != nil {
			users[i] = *u
		}
	}

	return users, nil
}

// Update implements UserRepository.Update
func (a *UserRepositoryAdapter) Update(
	ctx context.Context,
	user *entities.User,
) error {
	return a.genericRepo.Update(ctx, user)
}

// Delete implements UserRepository.Delete
func (a *UserRepositoryAdapter) Delete(ctx context.Context, id uint) error {
	return a.genericRepo.Delete(ctx, id)
}
