package auth

import (
	"context"
	"errors"
)

type contextKey string

const userIDContextKey contextKey = "user_id"
const userEmailContextKey contextKey = "user_email"

func SetUserInContext(ctx context.Context, userID string, email string) context.Context {
	ctx = context.WithValue(ctx, userIDContextKey, userID)
	ctx = context.WithValue(ctx, userEmailContextKey, email)
	return ctx
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, err := ctx.Value(userIDContextKey).(string)
	if !err {
		return "", errors.New("id not found in context")
	}
	return userID, nil
}

func GetUserEmailFromContext(ctx context.Context) (string, error) {
	email, err := ctx.Value(userEmailContextKey).(string)
	if !err {
		return "", errors.New("email not found in context")
	}
	return email, nil
}
