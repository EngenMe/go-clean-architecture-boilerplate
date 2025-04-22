package services

import (
	"context"

	"github.com/EngenMe/go-clean-architecture/domain/entities"
	"github.com/EngenMe/go-clean-architecture/infrastructure/utils"
	"github.com/EngenMe/go-clean-architecture/interfaces/repositories"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response to a successful login
type LoginResponse struct {
	Token string           `json:"token"`
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
) (*LoginResponse, error) {
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

	return &LoginResponse{
		Token: token,
		User:  user.ToDTO(),
	}, nil
}

// RegisterAuthService registers the auth service
func RegisterAuthService(userRepository repositories.UserRepository) *AuthService {
	return NewAuthService(userRepository)
}
