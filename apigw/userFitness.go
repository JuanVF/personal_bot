package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
)

type UserFitnessRouter struct {
}

// Register all the user fitness routes
func (userFitness UserFitnessRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s", userFitness.GetPrefix()), VerifyTokenMiddleware(userFitness.SetUserFitnessGoal)).Methods("POST")
}

// Returns the user fitness prefix
func (userFitness UserFitnessRouter) GetPrefix() string {
	return "/user/fitness"
}

// Create a new user fitness goal
func (userFitness UserFitnessRouter) SetUserFitnessGoal(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /user/fitness")

	var body *services.CreateUserFitnessBody = &services.CreateUserFitnessBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.AddUserFitnessGoal(body)

	writeResponse(w, resp)
}
