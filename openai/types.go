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

type GPT3Input struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             int     `json:"top_p"`
	N                int     `json:"n"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

type GPT3ChatInput struct {
	Model            string     `json:"model"`
	Messages         []GPT3Chat `json:"messages"`
	MaxTokens        int        `json:"max_tokens"`
	Temperature      float64    `json:"temperature"`
	TopP             int        `json:"top_p"`
	N                int        `json:"n"`
	FrequencyPenalty int        `json:"frequency_penalty"`
	PresencePenalty  int        `json:"presence_penalty"`
}

type GPT3Chat struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPT3Response struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"create"`
	Model   string       `json:"model"`
	Choices []GPT3Choice `json:"choices"`
}

type GPT3Choice struct {
	Text         string   `json:"text"`
	Index        int      `json:"index"`
	FinishReason string   `json:"finish_reason"`
	Message      GPT3Chat `json:"message"`
}

type GPT3Usage struct {
	PromptTokens     string `json:"prompt_tokens"`
	CompletionTokens string `json:"completion_tokens"`
	TotalTokens      string `json:"total_tokens"`
}
