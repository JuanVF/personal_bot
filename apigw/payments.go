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
	router.HandleFunc(fmt.Sprintf("%s", payment.GetPrefix()), VerifyTokenMiddleware(payment.GetPayments)).Methods("POST")
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
	}

	resp := services.GetPaymentsByTokenId(body)

	writeResponse(w, resp)
}
