package services

import (
	"fmt"
	"time"

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

// Converts a date from the standard RFC 822 to ISO String
func ConvertFROMRFC822toISOString(date string) string {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"
	t, err := time.Parse(layout, date)

	if err != nil {
		logger.Error("Service Helper", fmt.Sprintf("Error parsing the RFC822 Date[%s] to ISO String", date))
		return ""
	}

	return t.Format(time.RFC3339)
}
