package userservice

import (
	"context"
	"errors"
	"time"

	"github.com/maehiyu/tollo/internal/userservice/domain/user"
)

type CreateUserInput struct {
	ID      string
	Email   user.Email
	Name    string
	Profile user.Profile
}

type UpdateUserInput struct {
	ID      string
	Name    *string
	Profile *user.Profile
}

type Usecase interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*user.User, error)
	GetUserByID(ctx context.Context, id string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email user.Email) (*user.User, error)
	UpdateUser(ctx context.Context, input *UpdateUserInput) (*user.User, error)
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
	_, err := u.userRepo.FindByID(ctx, input.ID)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if !errors.Is(err, user.ErrNotFound) {
		return nil, err
	}

	newUser, err := user.NewUser(
		input.ID,
		input.Email,
		input.Name,
		input.Profile,
	)
	if err != nil {
		return nil, err
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

func (u *usecase) UpdateUser(ctx context.Context, input *UpdateUserInput) (*user.User, error) {
	targetUser, err := u.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		targetUser.Name = *input.Name
	}
	if input.Profile != nil {
		targetUser.Profile = *input.Profile
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
