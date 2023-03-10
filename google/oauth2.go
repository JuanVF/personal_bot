package google

import (
	"context"
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/monaco-io/request"
	"google.golang.org/api/idtoken"
)

// Given the user code, it will get the credentials information
func GetAccessToken(code string) (*OAuth2Response, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/token", config.Google.OAuthURL),
		Method: "POST",
		Query: map[string]string{
			"client_id":     config.Google.ClientId,
			"client_secret": config.Google.Secret,
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  config.Google.RedirectURI,
		},
	}

	var result *OAuth2Response = &OAuth2Response{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-OAuth2.0", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("Google-OAuth2.0", fmt.Sprintf("Invalid Code {%s}", code))
		return nil, fmt.Errorf("Invalid Code")
	}

	return result, nil
}

// Given a Refresh Token it will return the credentials information
func RefreshToken(refreshToken string) (*OAuth2Response, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/token", config.Google.OAuthURL),
		Method: "POST",
		Query: map[string]string{
			"client_id":     config.Google.ClientId,
			"client_secret": config.Google.Secret,
			"grant_type":    "refresh_token",
			"refresh_token": refreshToken,
		},
	}

	var result *OAuth2Response = &OAuth2Response{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-OAuth2.0", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("Google-OAuth2.0", fmt.Sprintf("Invalid Refresh Token {%s}", refreshToken))
		return nil, fmt.Errorf("Invalid Code")
	}

	return result, nil
}

// Given a Token it will verify if the token is valid
func IsValidToken(token string) bool {
	client := request.Client{
		URL:    fmt.Sprintf("%s/tokeninfo", config.Google.OAuthURL),
		Method: "GET",
		Query: map[string]string{
			"access_token": token,
		},
	}

	var result *VerifyTokenResponse = &VerifyTokenResponse{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-OAuth2.0", resp.Error().Error())
		return false
	}

	return resp.Response().StatusCode == 200
}

// Returns the payload from an ID Token
func GetPayloadFromIDToken(idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, config.Google.ClientId)

	if err != nil {
		common.GetLogger().Error("Google-OAuth2.0", err.Error())

		return nil, err
	}

	return payload, nil
}
