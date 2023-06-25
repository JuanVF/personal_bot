/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
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

// Register all the user health routes
func (userHealth UserHealthRouter) Handle() {
	router.HandleFunc(userHealth.GetPrefix(), VerifyTokenMiddleware(userHealth.SetUserHealth)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id_token}", userHealth.GetPrefix()), VerifyTokenMiddleware(userHealth.GetUserHealth)).Methods("GET")
}

// Returns the user health prefix
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
	urlParams := mux.Vars(r)
	idToken := urlParams["id_token"]

	if idToken == "" {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	logger.Log("APIGW", fmt.Sprintf("[GET] /user/health/%s", cropIdToken(idToken)))

	resp := services.GetUserHealth(idToken)

	writeResponse(w, resp)
}
