package apigw

import (
	"net/http"

	"github.com/JuanVF/personal_bot/google"
)

// Middleware to verify the access token send by the user
func VerifyTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		isValidToken := google.IsValidToken(token)

		if !isValidToken {
			http.Error(w, "Invalid Google OAuth 2.0 Access Token", http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	}
}
