package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound          = errors.New("user: not found")
	ErrEmailAlreadyExists = errors.New("user: email already exists")
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email Email) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}