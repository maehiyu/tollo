package user

import "errors"

var (
	ErrNotFound = errors.New("user: not founc")
	ErrEmailAlreadyExists = errors.New("user: email already exists")
	ErrInvalidID = errors.New("user: invalid id format")
	ErrInvalidEmail = errors.New("user: invalid email format")
  ErrNameRequired = errors.New("user: name is required") 
) 
