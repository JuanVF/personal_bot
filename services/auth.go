package services

import (
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
)

type GetTokenBody struct {
	Code string
}

// Service to retrieve an access token
func GetToken(body *GetTokenBody) *common.Response {
	data, err := google.GetAccessToken(body.Code)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusBadRequest)
	}

	return &common.Response{
		Status: http.StatusOK,
		Body:   data,
	}
}
