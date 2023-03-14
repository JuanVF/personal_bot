package services

import (
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/JuanVF/personal_bot/google"
	"github.com/JuanVF/personal_bot/repositories"
)

// Creates an user and its require data
func CreateUser(idToken, accessToken string) *common.Response {
	payload, err := google.GetPayloadFromIDToken(idToken)

	if err != nil {
		return common.GetErrorResponse("Invalid ID Token", http.StatusInternalServerError)
	}

	user, err := repositories.GetUserByGoogleMe(payload.Claims["email"].(string))

	if user != nil {
		return common.GetErrorResponse("This user is already registered in Personal Bot.", http.StatusInternalServerError)
	}

	createUser := &repositories.CreateUserBody{
		GoogleMe: payload.Claims["email"].(string),
		Name:     payload.Claims["given_name"].(string),
		LastName: payload.Claims["family_name"].(string),
	}

	err = repositories.CreateUser(createUser)

	if err != nil {
		return common.GetErrorResponse("There was en error creating your user. Please request help.", http.StatusInternalServerError)
	}

	bot := &repositories.CreateBotBody{
		UserId:      user.Id,
		LastGmailId: nil,
	}

	err = repositories.CreateBot(bot)

	if err != nil {
		return common.GetErrorResponse("There was en error creating your bot. Please request help.", http.StatusInternalServerError)
	}

	return nil
}
