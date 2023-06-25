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

type PaymentRouter struct {
}

// Register all the payment routes
func (payment PaymentRouter) Handle() {
	router.HandleFunc(fmt.Sprintf("%s/generate", payment.GetPrefix()), VerifyTokenMiddleware(payment.GeneratePayments)).Methods("POST")
	router.HandleFunc(payment.GetPrefix(), VerifyTokenMiddleware(payment.GetPayments)).Methods("POST")
}

// Returns the payment prefix
func (payment PaymentRouter) GetPrefix() string {
	return "/payments"
}

// Populates the latest payments in DB
func (payment PaymentRouter) GeneratePayments(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /payments/generate")

	var body *services.GeneratePaymentsBody = &services.GeneratePaymentsBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	body.BearerToken = r.Header.Get("Authorization")

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.GeneratePayments(body)

	writeResponse(w, resp)
}

// Returns the payments for a certain user using its token id
func (payment PaymentRouter) GetPayments(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /payments")

	var body *services.GetPaymentsBody = &services.GetPaymentsBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	resp := services.GetPaymentsByTokenId(body)

	writeResponse(w, resp)
}
