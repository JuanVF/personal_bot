package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
	"github.com/gorilla/mux"
)

type UserHealthRouter struct {
}

// Register all the payment routes
func (userHealth UserHealthRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s", userHealth.GetPrefix()), VerifyTokenMiddleware(userHealth.SetUserHealth)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id_token}", userHealth.GetPrefix()), VerifyTokenMiddleware(userHealth.GetUserHealth)).Methods("GET")
}

// Returns the payment prefix
func (userHealth UserHealthRouter) GetPrefix() string {
	return "/user/health"
}

// Set the user health data
func (userHealth UserHealthRouter) SetUserHealth(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /user/health")

	var body *services.AddUserHealthBody = &services.AddUserHealthBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.AddUserHealth(body)

	writeResponse(w, resp)
}

// Gets all the user health data
func (userHealth UserHealthRouter) GetUserHealth(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[GET] /user/health")

	urlParams := mux.Vars(r)
	idToken := urlParams["id_token"]

	if idToken == "" {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.GetUserHealth(idToken)

	writeResponse(w, resp)
}
