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

type GPT3Response struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"create"`
	Model   string       `json:"model"`
	Choices []GPT3Choice `json:"choices"`
}

type GPT3Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
}

type GPT3Usage struct {
	PromptTokens     string `json:"prompt_tokens"`
	CompletionTokens string `json:"completion_tokens"`
	TotalTokens      string `json:"total_tokens"`
}
