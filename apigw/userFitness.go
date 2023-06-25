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
)

type UserFitnessRouter struct {
}

// Register all the user fitness routes
func (userFitness UserFitnessRouter) Handle() {
	router.HandleFunc(userFitness.GetPrefix(), VerifyTokenMiddleware(userFitness.SetUserFitnessGoal)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/diet", userFitness.GetPrefix()), VerifyTokenMiddleware(userFitness.GenerateDietPlan)).Methods("POST")
}

// Returns the user fitness prefix
func (userFitness UserFitnessRouter) GetPrefix() string {
	return "/user/fitness"
}

// Generates a Diet Plan for the user based on the last walmart invoice
func (userFitness UserFitnessRouter) GenerateDietPlan(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", fmt.Sprintf("[POST] %s/diet", userFitness.GetPrefix()))

	token := r.Header.Get("Authorization")

	var body map[string]string = make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	idToken := body["id_token"]

	if idToken == "" {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := (&services.DietPlanService{}).GenerateDietPlan(idToken, token)

	writeResponse(w, resp)
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
