package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDevAuthMiddleware(t *testing.T) {
	t.Run("header exists", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := GetUserFromContext(r.Context())

			if err != nil {
				t.Errorf("context should have user_id")
			}
			if userID != "user-123" {
				t.Errorf("got %v, want user-123", userID)
			}
			w.WriteHeader(http.StatusOK)
		})
		handler := DevAuthMiddleware(nextHandler)

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-User-ID", "user-123")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)
	})
	t.Run("no header", func(t *testing.T) {
		nextCalled := false
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nextCalled = true
			_, err := GetUserFromContext(r.Context())

			if err != nil {
				t.Fatalf("GetUserFromContext() error = %v", err)
			}

			w.WriteHeader(http.StatusOK)
		})
		handler := DevAuthMiddleware(nextHandler)

		req := httptest.NewRequest("GET", "/text", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)
		if nextCalled != false {
			t.Errorf("next should not be called")
		}
		if w.Code != http.StatusUnauthorized {
			t.Errorf("status should be %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})
}
