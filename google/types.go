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

type OAuth2Response struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type GmailThreads struct {
	Messages           []GmailThreadIdentifier `json:"messages"`
	NextPageToken      string                  `json:"nextPageToken"`
	ResultSizeEstimate int                     `json:"resultSizeEstimate"`
}

type GmailThreadIdentifier struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
}

type GmailMessage struct {
	Id           string              `json:"id"`
	ThreadId     string              `json:"threadId"`
	LabelIds     []string            `json:"labelIds"`
	Snippet      string              `json:"snippet"`
	Payload      GmailMessagePayload `json:"payload"`
	SizeEstimate int                 `json:"sizeEstimate"`
	HistoryId    string              `json:"historyId"`
	InternalDate string              `json:"internalDate"`
}

type GmailMessagePayload struct {
	PartId   string                `json:"partId"`
	MimeType string                `json:"mimeType"`
	Filename string                `json:"filename"`
	Headers  []GmailMessageHeader  `json:"headers"`
	Body     GmailMessageBody      `json:"body"`
	Parts    []GmailMessagePayload `json:"parts"`
}

type GmailMessageHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GmailMessageBody struct {
	Size         int    `json:"size"`
	Data         string `json:"data"`
	AttachmentId string `json:"attachmentId"`
}

type VerifyTokenResponse struct {
	Azp        string `json:"azp"`
	Aud        string `json:"aud"`
	Scope      string `json:"scope"`
	Exp        string `json:"exp"`
	ExpiresIn  string `json:"expires_in"`
	AccessType string `json:"access_type"`
}

type IDTokenInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
}
