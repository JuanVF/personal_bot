package services

import (
	"fmt"

	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/repositories"
)

// Given an ID Token it will return the user
func getUserByIDToken(idToken string) (*repositories.User, error) {
	payload, err := google.GetPayloadFromIDToken(idToken)

	if err != nil {
		return nil, fmt.Errorf("Invalid ID Token")
	}

	user, err := repositories.GetUserByGoogleMe(payload.Claims["email"].(string))

	if err != nil {
		return nil, fmt.Errorf("This user is not registered in Personal Bot")
	}

	return user, nil
}
