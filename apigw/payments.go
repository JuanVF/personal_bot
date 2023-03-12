package apigw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanVF/personal_bot/services"
)

var paymentPrefix string = "/payments"

// Register all the payment routes
func HandlePaymentRoutes() {
	router.HandleFunc(fmt.Sprintf("%s/generate", paymentPrefix), VerifyTokenMiddleware(GeneratePayments)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s", paymentPrefix), VerifyTokenMiddleware(GetPayments)).Methods("POST")
}

// Populates the latest payments in DB
func GeneratePayments(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /payments/generate")

	var body *services.GeneratePaymentsBody = &services.GeneratePaymentsBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	body.BearerToken = r.Header.Get("Authorization")

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
	}

	resp := services.GeneratePayments(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp.Body)
}

// Returns the payments for a certain user using its token id
func GetPayments(w http.ResponseWriter, r *http.Request) {
	logger.Log("APIGW", "[POST] /payments")

	var body *services.GetPaymentsBody = &services.GetPaymentsBody{}

	err := json.NewDecoder(r.Body).Decode(body)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
	}

	resp := services.GetPaymentsByTokenId(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp.Body)
}
