package openai

import (
	"fmt"

	"github.com/JuanVF/personal_bot/common"
	"github.com/monaco-io/request"
)

// Uses OpenAI API to query a Fined Tunes Model for labeling
func Complete(tokens string) (*GPT3Response, error) {
	completeParams := config.OpenAI.CompleteParams

	client := request.Client{
		URL:    fmt.Sprintf("%s/v1/completions", config.OpenAI.OpenAIAPI),
		Method: "POST",
		JSON: GPT3Input{
			Model:            config.OpenAI.FinedTunedModel,
			Prompt:           tokens,
			MaxTokens:        completeParams.MaxTokens,
			Temperature:      completeParams.Temperature,
			TopP:             completeParams.TopP,
			N:                completeParams.N,
			FrequencyPenalty: completeParams.FrequencyPenalty,
			PresencePenalty:  completeParams.PresencePenalty,
		},
		Header: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", config.OpenAI.SecretKey),
		},
	}

	var result *GPT3Response = &GPT3Response{}

	resp := client.Send().Scan(&result)

	if !resp.OK() {
		common.GetLogger().Error("OpenAI", resp.Error().Error())
		return nil, resp.Error()
	}

	if resp.Response().StatusCode != 200 {
		common.GetLogger().Error("OpenAI", fmt.Sprintf("Error while completing"))
		return nil, fmt.Errorf("Error while completing")
	}

	return result, nil
}
