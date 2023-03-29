package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
)

type UserHealthRouter struct {
}

// Register all the payment routes
func (userHealth UserHealthRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s", userHealth.GetPrefix()), VerifyTokenMiddleware(userHealth.SetUserHealth)).Methods("POST")
}

// Returns the payment prefix
func (userHealth UserHealthRouter) GetPrefix() string {
	return "/user/health"
}

// Populates the latest payments in DB
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
