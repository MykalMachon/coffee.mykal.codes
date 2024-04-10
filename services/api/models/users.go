package models

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    int32
}

type UserService struct{}
