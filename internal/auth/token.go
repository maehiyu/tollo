package auth

import (
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
}

type TokenGenerator interface {
	GenerateToke(userID string, duration time.Duration) (string, error)
}

type TokenVerifier interface {
	VerifyToken(tokenString string) (*Claims, error)
}
