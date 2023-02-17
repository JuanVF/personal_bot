package google

import (
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/monaco-io/request"
)

var config *common.Configuration = common.GetConfig()

// Given the user code, it will get the credentials information
func GetAccessToken(code string) (*OAuth2Response, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/token", config.Gmail.BaseURL),
		Method: "POST",
		Query: map[string]string{
			"client_id":     config.Gmail.ClientId,
			"client_secret": config.Gmail.Secret,
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  config.Gmail.RedirectURI,
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
		URL:    fmt.Sprintf("%s/token", config.Gmail.BaseURL),
		Method: "POST",
		Query: map[string]string{
			"client_id":     config.Gmail.ClientId,
			"client_secret": config.Gmail.Secret,
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
