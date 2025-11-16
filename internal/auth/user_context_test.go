package auth

import (
	"context"
	"testing"
)

func TestSetUserOnContext(t *testing.T) {
	tests := []struct {
		name    string
		input   context.Context
		want    string
		wantErr bool
	}{
		{
			name:    "user exists",
			input:   SetUserInContext(context.Background(), "user-123"),
			want:    "user-123",
			wantErr: false,
		},
		{
			name:    "no user",
			input:   context.Background(),
			want:    "",
			wantErr: true,
		}, {
			name:    "wrong type",
			input:   context.WithValue(context.Background(), userContextKey, 123),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserFromContext(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFromContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
