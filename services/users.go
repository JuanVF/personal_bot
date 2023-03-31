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
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/repositories"
)

// Creates an user and its require data
func CreateUser(idToken, accessToken string) *common.Response {
	user, err := createUserByIdToken(idToken)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	err = createBotByUser(accessToken, user)

	if err != nil {
		return common.GetErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	return common.GetSuccessResponse(map[string]string{"Message": "User Created"})
}

// Creates an user by its ID Token
func createUserByIdToken(idToken string) (*repositories.User, error) {
	payload, err := google.GetPayloadFromIDToken(idToken)

	if err != nil {
		return nil, fmt.Errorf("Invalid ID Token")
	}

	user, err := repositories.GetUserByGoogleMe(payload.Claims["email"].(string))

	if user != nil {
		return nil, fmt.Errorf("This user is already registered in Personal Bot.")
	}

	createUser := &repositories.CreateUserBody{
		GoogleMe: payload.Claims["email"].(string),
		Name:     payload.Claims["given_name"].(string),
		LastName: payload.Claims["family_name"].(string),
	}

	userId, err := repositories.CreateUser(createUser)

	if err != nil {
		return nil, fmt.Errorf("There was en error creating your user. Please request help.")
	}

	return &repositories.User{
		Id:       userId,
		GoogleMe: payload.Claims["email"].(string),
		Name:     payload.Claims["given_name"].(string),
		LastName: payload.Claims["family_name"].(string),
	}, nil
}

// Creates a bot for an user
func createBotByUser(accessToken string, user *repositories.User) error {
	messages, err := google.GetGmailMessageList(user.GoogleMe, googleQuery, accessToken)

	if err != nil {
		logger.Error("User Service - Create User", err.Error())
		return fmt.Errorf("There was en error requesting your mails. Please request help.")
	}

	lastGmailId := ""

	if len(messages.Messages) > 0 {
		lastGmailId = messages.Messages[0].Id
	}

	bot := &repositories.CreateBotBody{
		UserId:      user.Id,
		LastGmailId: &lastGmailId,
	}

	err = repositories.CreateBot(bot)

	if err != nil {
		return fmt.Errorf("There was en error creating your bot. Please request help.")
	}

	return nil
}

// Returns all the personal data stored in personal bot
func GetPersonalData(idToken string) *common.Response {
	user, err := getUserByIDToken(idToken)

	if err != nil {
		return common.GetErrorResponse("Invalid ID Token", http.StatusUnauthorized)
	}

	bot, err := repositories.GetBotByUserId(user.Id)

	if err != nil {
		return common.GetErrorResponse("There was en error requesting your bot. Please request help.", http.StatusInternalServerError)
	}

	healthData, err := repositories.GetUserHealthByUserId(user.Id)

	if err != nil {
		return common.GetErrorResponse("There was en error requesting your health data. Please request help.", http.StatusInternalServerError)
	}

	fitnessGoals, err := repositories.GetFitnessGoalsByUser(user.Id)

	if err != nil {
		return common.GetErrorResponse("There was en error requesting your fitness goals. Please request help.", http.StatusInternalServerError)
	}

	return common.GetSuccessResponse(map[string]interface{}{
		"user":          user,
		"bot":           bot,
		"health_data":   healthData,
		"fitness_goals": fitnessGoals,
	})
}
