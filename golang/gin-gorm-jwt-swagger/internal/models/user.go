package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is the struct to manipulate user operations.
type User struct {
	ID        *uint     `json:"id,omitempty" binding:"excluded_with_all" gorm:"primarykey"`
	Name      string    `json:"name,omitempty" gorm:"default:unnamed"`
	Username  *string   `json:"username,omitempty" binding:"required" gorm:"unique;not null"`
	Email     *string   `json:"email,omitempty" binding:"required,email" gorm:"unique;not null"`
	Password  *string   `json:"password,omitempty" binding:"required,min=8" gorm:"not null"`
	CreatedAt time.Time `binding:"excluded_with_all"`
	UpdatedAt time.Time `binding:"excluded_with_all"`
}

// HashPassword hash the password of entity.
func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(*user.Password),
		bcrypt.DefaultCost,
	)
	*user.Password = string(hashedPassword)
	return err
}

// IsValidPassword validate the passing password.
func (user *User) IsValidPassword(passwordToCompare string) bool {
	// return bcrypt.CompareHashAndPassword([]byte(passwordToCompare), []byte(testUser.Password)) == nil
	err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(passwordToCompare),
	)
	return err == nil
}
