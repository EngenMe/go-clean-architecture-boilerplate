package entities

import (
	"time"
)

// User represents a user entity in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // Password is never exposed in JSON responses
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the ID of the user
// Note: This is a value receiver to implement the Entity interface
func (u User) GetID() uint {
	return u.ID
}

// SetID sets the ID of the user
// Note: This is a value receiver to implement the Entity interface
func (u User) SetID(id uint) {
	u.ID = id
}

// TableName specifies the table name for the User entity
func (User) TableName() string {
	return "users"
}

// UserDTO is a data transfer object for User entity
type UserDTO struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
