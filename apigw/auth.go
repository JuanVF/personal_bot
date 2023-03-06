package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/services"
)

var authPrefix string = "/auth"

// Register all the auth routes
func HandleAuthRoutes() {
	router.HandleFunc(fmt.Sprintf("%s/token", authPrefix), GetToken).Methods("GET")
}

// Get Token End Point will request an access token with a code
func GetToken(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	code := values.Get("code")

	var body *services.GetTokenBody = &services.GetTokenBody{
		Code: code,
	}

	logger.LogObject("APIGW - Get Token", *body)

	resp := services.GetToken(body)

	if resp.Status != 200 {
		http.Error(w, "", resp.Status)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Body)
}

// Middleware to verify the access token send by the user
func VerifyTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Header := w.Header()
		token := Header.Get("Authorization")

		isValidToken := google.IsValidToken(token)

		if !isValidToken {
			http.Error(w, "Invalid Google OAuth 2.0 Access Token", http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	}
}