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
