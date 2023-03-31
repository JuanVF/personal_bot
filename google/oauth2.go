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
