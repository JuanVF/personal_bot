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
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/monaco-io/request"
)

// Returns the user gmails 100 latest messages filtered by the google q param
func GetGmailMessageList(me, q, bearerToken string) (*GmailThreads, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/gmail/v1/users/%s/messages", config.Google.GmailURL, me),
		Method: "GET",
		Query: map[string]string{
			"q": q,
		},
		Bearer: bearerToken,
	}

	var result *GmailThreads = &GmailThreads{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-Gmail", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("Google-Gmail", fmt.Sprintf("<%s> for user <%s>", resp.Response().Status, me))
		return nil, fmt.Errorf("Invalid Code")
	}

	return result, nil
}

// Retrieves a specific thread or message from gmail
func GetGmailMessage(me, threadId, bearerToken string) (*GmailMessage, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/gmail/v1/users/%s/messages/%s", config.Google.GmailURL, me, threadId),
		Method: "GET",
		Bearer: bearerToken,
	}

	var result *GmailMessage = &GmailMessage{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-Gmail", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("Google-Gmail", fmt.Sprintf("<%s> for user <%s> and thredId <%s>", resp.Response().Status, me, threadId))
		return nil, fmt.Errorf("Invalid Code")
	}

	return result, nil
}

// Retrieves a specific body from a message
func RetrieveFile(me, threadId, attachmentId, bearerToken string) (*GmailMessageBody, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/gmail/v1/users/%s/messages/%s/attachments/%s", config.Google.GmailURL, me, threadId, attachmentId),
		Method: "GET",
		Bearer: bearerToken,
	}

	var result *GmailMessageBody = &GmailMessageBody{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("Google-Gmail", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("Google-Gmail", fmt.Sprintf("<%s> for user <%s> and thredId <%s> and attachmentId <%s>", resp.Response().Status, me, threadId, attachmentId))
		return nil, fmt.Errorf("Invalid Code")
	}

	return result, nil

}
