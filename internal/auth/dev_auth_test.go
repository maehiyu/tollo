package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDevAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		useridheader   string
		emailheader    string
		wantstatus     int
		wantuserid     string
		wantemail      string
		expectnextcall bool
	}{
		{
			name:           "both headers exist",
			useridheader:   "user-123",
			emailheader:    "test@example.com",
			wantstatus:     http.StatusOK,
			wantuserid:     "user-123",
			wantemail:      "test@example.com",
			expectnextcall: true,
		},
		{
			name:           "only user-id header",
			useridheader:   "user-123",
			emailheader:    "",
			wantstatus:     http.StatusUnauthorized,
			wantuserid:     "",
			wantemail:      "",
			expectnextcall: false,
		},
		{
			name:           "only email header",
			useridheader:   "",
			emailheader:    "test@example.com",
			wantstatus:     http.StatusUnauthorized,
			wantuserid:     "",
			wantemail:      "",
			expectnextcall: false,
		},
		{
			name:           "no headers",
			useridheader:   "",
			emailheader:    "",
			wantstatus:     http.StatusUnauthorized,
			wantuserid:     "",
			wantemail:      "",
			expectnextcall: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextCalled := false
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true

				// expectnextcallがfalseなのに呼ばれた場合はエラー
				if !tt.expectnextcall {
					t.Error("next handler should not be called")
				}

				gotID, IDErr := GetUserIDFromContext(r.Context())
				gotEmail, emailErr := GetUserEmailFromContext(r.Context())

				// エラーチェック
				if tt.expectnextcall {
					if IDErr != nil {
						t.Errorf("GetUserIDFromContext() should not error, got %v", IDErr)
					}
					if emailErr != nil {
						t.Errorf("GetUserEmailFromContext() should not error, got %v", emailErr)
					}
				}

				if gotID != tt.wantuserid {
					t.Errorf("gotID %v, want %v", gotID, tt.wantuserid)
				}
				if gotEmail != tt.wantemail {
					t.Errorf("gotEmail %v, want %v", gotEmail, tt.wantemail)
				}
				w.WriteHeader(http.StatusOK)
			})
			handler := DevAuthMiddleware(nextHandler)

			req := httptest.NewRequest("GET", "/test", nil)
			if tt.useridheader != "" {
				req.Header.Set("X-User-ID", tt.useridheader)
			}
			if tt.emailheader != "" {
				req.Header.Set("X-User-Email", tt.emailheader)
			}
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			if w.Code != tt.wantstatus {
				t.Errorf("status should be %d, got %d", tt.wantstatus, w.Code)
			}

			if nextCalled != tt.expectnextcall {
				t.Errorf("nextCalled = %v, want %v", nextCalled, tt.expectnextcall)
			}
		})
	}
}
