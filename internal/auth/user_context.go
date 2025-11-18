package auth

import (
	"context"
)

type contextKey string

const userIDContextKey contextKey = "user_id"
const userEmailContextKey contextKey = "user_email"

func SetUserInContext(ctx context.Context, userID string, email string) context.Context {
	ctx = context.WithValue(ctx, userIDContextKey, userID)
	ctx = context.WithValue(ctx, userEmailContextKey, email)
	return ctx
}

func MustGetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(userIDContextKey).(string)
	if !ok {
		panic("User ID not in context - middleware misconfigured")
	}

	return userID
}

func MustGetUserEmailFromContext(ctx context.Context) string {
	email, ok := ctx.Value(userEmailContextKey).(string)
	if !ok {
		panic("Email not in context - middleware misconfigured")
	}
	return email
}
