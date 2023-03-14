package apigw

import (
	"encoding/json"
	"net/http"

	"github.com/JuanVF/personal_bot/common"
)

// Writes the response for a common JSON response
func writeResponse(w http.ResponseWriter, response *common.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Body)
}