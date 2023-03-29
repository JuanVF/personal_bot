package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
	"github.com/gorilla/mux"
)

type UserRouter struct {
}

// Router Handler for user
func (u UserRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s", u.GetPrefix()), VerifyTokenMiddleware(u.CreateUser)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id_token}", u.GetPrefix()), VerifyTokenMiddleware(u.GetData)).Methods("GET")
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

// Creates an user using the ID Token
func (u UserRouter) GetData(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	idToken := urlParams["id_token"]

	if idToken == "" {
		http.Error(w, "Invalid ID Token", http.StatusBadRequest)
		return
	}

	logger.Log("APIGW", fmt.Sprintf("[GET] %s/%s", u.GetPrefix(), cropIdToken(idToken)))

	resp := services.GetPersonalData(idToken)

	writeResponse(w, resp)
}
