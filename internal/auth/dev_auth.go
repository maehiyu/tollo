package auth

import "net/http"

func DevAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		if userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := SetUserInContext(r.Context(), userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
