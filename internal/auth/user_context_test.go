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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID := MustGetUserIDFromContext(tt.input)
			gotEmail := MustGetUserEmailFromContext(tt.input)

			if gotID != tt.wantID {
				t.Errorf("got %v, want %v", gotID, tt.wantID)
			}
			if gotEmail != tt.wantEmail {
				t.Errorf("got %v, want %v", gotEmail, tt.wantEmail)
			}
		})
	}
}
