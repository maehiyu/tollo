package user

import "errors"

var (
	ErrNotFound           = errors.New("user: not found")
	ErrEmailAlreadyExists = errors.New("user: email already exists")
	ErrInvalidID          = errors.New("user: invalid id format")
	ErrEmptyID            = errors.New("user: id cannot be empty")
	ErrEmptyName          = errors.New("user: name cannot be empty")
	ErrNilProfile         = errors.New("user: profile cannot be nil")
	ErrInvalidEmail       = errors.New("user: invalid email format")
	ErrNameRequired       = errors.New("user: name is required")
)
