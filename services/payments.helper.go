package services

import (
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/repositories"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

type PaymentData struct {
	Body     string
	Currency *repositories.Currency
	Amount   float64
}

func GetHeader(headers *[]google.GmailMessageHeader, headerName string) (string, error) {
	for _, header := range *headers {
		if header.Name == headerName {
			return header.Value, nil
		}
	}

	return "", fmt.Errorf("Header not found")
}

// Will return the body from the payload, since it is a recursive struct we will use a recursive way
// Since it has low recursive rate is OK to use it
func GetBodyFromPayload(payload *google.GmailMessagePayload) (string, error) {
	if payload == nil {
		return "", fmt.Errorf("Body not found")
	}

	if payload.Body.Size == 0 {
		return GetBodyFromPayload(&payload.Parts[0])
	}

	decoded, err := base64url.Decode(payload.Body.Data)

	if err != nil {
		common.GetLogger().Error("Payment Service Helper", err.Error())
		return "", fmt.Errorf("Error while decoding Thread with PartId [%s]", payload.PartId)
	}

	return string(decoded), nil
}
