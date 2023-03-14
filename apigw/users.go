package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
)

type UserRouter struct {
}

// Router Handler for user
func (u UserRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s", u.GetPrefix()), VerifyTokenMiddleware(u.CreateUser)).Methods("POST")
}

// Router Handler for user
func (u UserRouter) GetPrefix() string {
	return "/users"
}

// Creates an user using the ID Token
func (u UserRouter) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /users")

	body := make(map[string]interface{})
	token := r.Header.Get("Authorization")

	// Decode the request body into the data map
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.CreateUser(body["id_token"].(string), token)

	writeResponse(w, resp)
}
