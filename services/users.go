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
