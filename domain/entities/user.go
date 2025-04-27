package entities

import (
	"time"
)

// User represents a user entity in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null" example:"user@example.com"`
	Password  string    `json:"-" gorm:"not null"` // Password is never exposed in JSON responses
	FirstName string    `json:"firstName" gorm:"not null" example:"John"`
	LastName  string    `json:"lastName" gorm:"not null" example:"Doe"`
	CreatedAt time.Time `json:"createdAt" example:"2025-04-27T12:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-04-27T12:00:00Z"`
}

// GetID returns the ID of the user
func (u User) GetID() uint {
	return u.ID
}

// SetID sets the ID of the user
func (u User) SetID(id uint) {
	u.ID = id
}

// TableName specifies the table name for the User entity
func (User) TableName() string {
	return "users"
}

// UserDTO is a data transfer object for User entity
type UserDTO struct {
	ID        uint      `json:"id" example:"1"`
	Email     string    `json:"email" example:"user@example.com"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	CreatedAt time.Time `json:"createdAt" example:"2025-04-27T12:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-04-27T12:00:00Z"`
}

// ToDTO converts a User entity to UserDTO
func (u User) ToDTO() UserDTO {
	return UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
