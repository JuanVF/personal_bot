package google

import (
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/monaco-io/request"
)

// Returns the user gmails 100 latest messages
func GetGmailMessageList(me string, bearerToken string) (*GmailThreads, error) {
	client := request.Client{
		URL:    fmt.Sprintf("%s/gmail/v1/users/%s/messages", config.Google.GmailURL, me),
		Method: "GET",
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
