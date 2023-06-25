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

type UserRouter struct {
}

// Router Handler for user
func (u UserRouter) Handle() {
	router.HandleFunc(u.GetPrefix(), VerifyTokenMiddleware(u.CreateUser)).Methods("POST")
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

	cuBody := &services.CreateUserBody{
		IdToken:       body["id_token"].(string),
		AccessToken:   token,
		ActivityLevel: body["activity_level"].(string),
	}

	resp := services.CreateUser(cuBody)

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
