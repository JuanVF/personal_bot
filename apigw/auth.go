package apigw

import (
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
)

type AuthRoute struct {
}

// Register all the auth routes
func (auth AuthRoute) Handle() {
	router.HandleFunc(fmt.Sprintf("%s/token", auth.GetPrefix()), auth.GetToken).Methods("GET")
}

// Returns the prefix for this Handler
func (auth AuthRoute) GetPrefix() string {
	return "/auth"
}

// Get Token End Point will request an access token with a code
func (auth AuthRoute) GetToken(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /auth")

	values := r.URL.Query()
	code := values.Get("code")

	var body *services.GetTokenBody = &services.GetTokenBody{
		Code: code,
	}

	resp := services.GetToken(body)

	writeResponse(w, resp)
}
