package services

import (
	"context"
	"fmt"

	"github.com/EngenMe/go-clean-architecture/application/commands"
	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"github.com/mehdihadeli/go-mediatr"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// SignUpRequest represents sign-up data
type SignUpRequest struct {
	Email     string `json:"email" binding:"required,email" example:"user@example.com"`
	Password  string `json:"password" binding:"required,min=6" example:"password123"`
	FirstName string `json:"firstName" binding:"required" example:"John"`
	LastName  string `json:"lastName" binding:"required" example:"Doe"`
}

// AuthResponse represents the response to a successful authentication
type AuthResponse struct {
	Token string           `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  entities.UserDTO `json:"user"`
}

// AuthService provides authentication functionality
type AuthService struct {
	userRepository repositories.UserRepository
}

// NewAuthService creates a new authentication service
func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

// Login authenticates a user and generates a JWT token
func (s *AuthService) Login(
	ctx context.Context,
	request LoginRequest,
) (*AuthResponse, error) {
	// Find user by email
	user, err := s.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrUnauthorized
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return nil, utils.ErrUnauthorized
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user.ToDTO(),
	}, nil
}

// SignUp registers a new user and generates a JWT token
func (s *AuthService) SignUp(
	ctx context.Context,
	request SignUpRequest,
) (*AuthResponse, error) {
	// Create user command
	command := commands.CreateUserCommand{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	// Execute command via mediatr
	result, err := mediatr.Send[commands.CreateUserCommand, *entities.UserDTO](
		ctx,
		command,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Find the newly created user to get full entity with ID
	user, err := s.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &AuthResponse{
		Token: token,
		User:  *result,
	}, nil
}

// RegisterAuthService registers the auth service
func RegisterAuthService(userRepository repositories.UserRepository) *AuthService {
	return NewAuthService(userRepository)
}
