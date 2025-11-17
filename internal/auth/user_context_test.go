package auth

import (
	"context"
	"testing"
)

func TestSetUserOnContext(t *testing.T) {
	tests := []struct {
		name      string
		input     context.Context
		wantID    string
		wantEmail string
		wantErr   bool
	}{
		{
			name:      "user exists",
			input:     SetUserInContext(context.Background(), "user-123", "test@sample.com"),
			wantID:    "user-123",
			wantEmail: "test@sample.com",
			wantErr:   false,
		},
		{
			name:      "no user",
			input:     context.Background(),
			wantID:    "",
			wantEmail: "",
			wantErr:   true,
		},
		{
			name:      "wrong id type",
			input:     context.WithValue(context.WithValue(context.Background(), userIDContextKey, 123), userEmailContextKey, "test@sample.com"),
			wantID:    "",
			wantEmail: "test@sample.com",
			wantErr:   true,
		},
		{
			name:      "wrong email type",
			input:     context.WithValue(context.WithValue(context.Background(), userIDContextKey, "user-123"), userEmailContextKey, 123),
			wantID:    "user-123",
			wantEmail: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, idErr := GetUserIDFromContext(tt.input)
			gotEmail, EmailErr := GetUserEmailFromContext(tt.input)

			if (idErr != nil || EmailErr != nil) != tt.wantErr {
				t.Errorf("error = %v / %v, wantErr %v", idErr, EmailErr, tt.wantErr)
			}
			if gotID != tt.wantID {
				t.Errorf("got %v, want %v", gotID, tt.wantID)
			}
			if gotEmail != tt.wantEmail {
				t.Errorf("got %v, want %v", gotEmail, tt.wantEmail)
			}
		})
	}
}
