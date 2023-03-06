package services

import (
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
)

type GetTokenBody struct {
	Code string
}

func GetToken(body *GetTokenBody) *common.Response {
	var response *common.Response = &common.Response{
		Status: http.StatusOK,
	}

	data, err := google.GetAccessToken(body.Code)

	if err != nil {
		response.Status = http.StatusBadRequest
		response.Body = &common.ErrorResponse{
			Message: err.Error(),
		}

		return response
	}

	response.Body = data

	return response
}
