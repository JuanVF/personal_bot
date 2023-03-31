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
