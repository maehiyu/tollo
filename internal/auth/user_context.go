package auth

import (
	"context"
	"errors"
)

type contextKey string

const userContextKey contextKey = "user_id"

func SetUserInContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userContextKey, userID)
}

func GetUserFromContext(ctx context.Context) (string, error) {
	userID, err := ctx.Value(userContextKey).(string)
	if !err {
		return "", errors.New("user not found in context")
	}
	return userID, nil
}
