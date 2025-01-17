package domain

import "context"

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
