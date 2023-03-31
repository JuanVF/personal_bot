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

// Uses OpenAI API to query a Fined Tunes Model for labeling
func Chat(chat *[]GPT3Chat) (*GPT3Response, error) {
	completeParams := config.OpenAI.CompleteParams

	client := request.Client{
		URL:    fmt.Sprintf("%s/v1/chat/completions", config.OpenAI.OpenAIAPI),
		Method: "POST",
		JSON: GPT3ChatInput{
			Model:            "gpt-3.5-turbo",
			MaxTokens:        500,
			Messages:         *chat,
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
		common.GetLogger().LogObject("OpenAI", resp.Response().Status)

		return nil, fmt.Errorf("Error while completing")
	}

	return result, nil
}
