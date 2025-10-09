package userservice

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/maehiyu/tollo/internal/userservice/domain/user"
)


type CreateUserInput struct {
	Name        string
	Email       user.Email
	Description string
}

type UpdateUserInput struct {
	Name        *string
	Description *string
}


type Usecase interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*user.User, error)
	GetUserByID(ctx context.Context, id string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email user.Email) (*user.User, error)
	UpdateUser(ctx context.Context, id string, input *UpdateUserInput) (*user.User, error)
	DeleteUser(ctx context.Context, id string) error
}


type usecase struct {
	userRepo user.UserRepository
}

func NewUsecase(userRepo user.UserRepository) Usecase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (u *usecase) CreateUser(ctx context.Context, input *CreateUserInput) (*user.User, error) {
	_, err := u.userRepo.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, user.ErrNotFound) { 
		return nil, err
	}
	if err == nil {
		return nil, user.ErrEmailAlreadyExists 
	}

	now := time.Now()
	newUser := &user.User{
		ID:          uuid.NewString(),
		Name:        input.Name,
		Email:       input.Email,
		Description: input.Description,
		Type:        user.GeneralUser,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := u.userRepo.Save(ctx, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *usecase) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	foundUser, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *usecase) GetUserByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	foundUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *usecase) UpdateUser(ctx context.Context, id string, input *UpdateUserInput) (*user.User, error) {
	targetUser, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		targetUser.Name = *input.Name
	}
	if input.Description != nil {
		targetUser.Description = *input.Description
	}
	targetUser.UpdatedAt = time.Now()

	if err := u.userRepo.Save(ctx, targetUser); err != nil {
		return nil, err
	}
	return targetUser, nil
}

func (u *usecase) DeleteUser(ctx context.Context, id string) error {
	_, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return u.userRepo.Delete(ctx, id)
}