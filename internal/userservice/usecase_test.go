package userservice

import (
	"context"
	"testing"

	"github.com/maehiyu/tollo/internal/adapter/repository"
	"github.com/maehiyu/tollo/internal/userservice/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	uc := NewUsecase(mockRepo)

	t.Run("create user", func(t *testing.T) {
		ctx := context.Background()
		email, err := user.NewEmail("test@sample.com")
		createdUser, err := uc.CreateUser(ctx, &CreateUserInput{
			ID:      "user-123",
			Email:   email,
			Name:    "test",
			Profile: &user.GeneralProfile{Points: 1, Introduction: "hello"},
		})
		assert.NoError(t, err)

		assert.Equal(t, "test", createdUser.Name)
		assert.Equal(t, "user-123", createdUser.ID)
		assert.Equal(t, user.Email("test@sample.com"), createdUser.Email)
	})
	t.Run("user already exist", func(t *testing.T) {
		ctx := context.Background()
		email, err := user.NewEmail("test@sample.com")
		createUser, err := uc.CreateUser(ctx, &CreateUserInput{
			ID:      "user-123",
			Email:   email,
			Name:    "test",
			Profile: &user.GeneralProfile{Points: 1, Introduction: "hello"},
		})
		assert.Error(t, err)
		assert.Nil(t, createUser)
	})
}
func TestUpdateUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	uc := NewUsecase(mockRepo)

	t.Run("update user", func(t *testing.T) {
		ctx := context.Background()

		// まずユーザーを作成
		email, _ := user.NewEmail("test@sample.com")
		createdUser, err := uc.CreateUser(ctx, &CreateUserInput{
			ID:      "user-123",
			Email:   email,
			Name:    "original",
			Profile: &user.GeneralProfile{Points: 1, Introduction: "hello"},
		})
		assert.NoError(t, err)

		// 次に更新
		newName := "changed"
		updatedUser, err := uc.UpdateUser(ctx, &UpdateUserInput{
			ID:      createdUser.ID,
			Name:    &newName,
			Profile: nil,
		})
		assert.NoError(t, err)

		assert.Equal(t, "changed", updatedUser.Name)
		assert.Equal(t, "user-123", updatedUser.ID)
	})
}
