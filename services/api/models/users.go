package models

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Posts        []Post
}

type UserService struct {
	DB *gorm.DB
}

func (us UserService) CreateUser(name string, email string, password string) (*User, error) {
	// TODO: check if signup flag is enabled
	email = strings.ToLower(email) // lowercase email always
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	passwordHash := string(hashedPasswordBytes)

	user := User{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
	}

	result := us.DB.Create(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("create user: %w", result.Error)
	}

	return &User{}, nil
}

func (us UserService) Authenticate(email string, password string) (*User, error) {
	email = strings.ToLower(email)

	// get user with this email
	user := User{}
	result := us.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cannot authenticate as no user exists with this email: %w", result.Error)
		} else {
			return nil, fmt.Errorf("authenticate user: %w", result.Error)
		}
	}

	// compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate %w", err)
	}

	return &user, nil
}

func (us UserService) CompletePasswordReset(token string, newPassword string) {
	// lookup token in "password reset" in database
	// ensure it is not expired and has not yet been used
	// if it is valid, allow the user to replace their password
}

func (us UserService) RequestPasswordReset() {
	// create "password reset" item in database
	// send an email to the user asking to reset their password
}
