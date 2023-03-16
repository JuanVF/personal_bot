package common

import "database/sql"

type EnvironmentConfig struct {
	Development Configuration `yaml:"development"`
	Container   Configuration `yaml:"container"`
}

type Configuration struct {
	Google        GoogleConf        `yaml:"google"`
	Bot           BotConf           `yaml:"bot"`
	PersonalBotDB PersonalBotDBConf `yaml:"personal_bot_db"`
	OpenAI        OpenAIConfig      `yaml:"open_ai"`
}

type OpenAIConfig struct {
	OpenAIAPI       string               `yaml:"open_ai_api"`
	SecretKey       string               `yaml:"secret_key"`
	FinedTunedModel string               `yaml:"fined_tuned_model"`
	CompleteParams  OpenAICompleteParams `yaml:"complete_params"`
}

type OpenAICompleteParams struct {
	MaxTokens        int     `yaml:"max_tokens"`
	Temperature      float64 `yaml:"temperature"`
	TopP             int     `yaml:"top_p"`
	N                int     `yaml:"n"`
	FrequencyPenalty int     `yaml:"frequency_penalty"`
	PresencePenalty  int     `yaml:"presence_penalty"`
}

type PersonalBotDBConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"ssl_mode"`
}

type BotConf struct {
	ReadFrom string `yaml:"read_from"`
}

type GoogleConf struct {
	OAuthURL    string `yaml:"oauth_url"`
	GmailURL    string `yaml:"gmail_url"`
	ClientId    string `yaml:"client_id"`
	Secret      string `yaml:"secret"`
	RedirectURI string `yaml:"redirect_uri"`
}

type Logger struct {
	headline string
}

type DB struct {
	connection *sql.DB
}

type Response struct {
	Status int
	Body   any
}

type ErrorResponse struct {
	Message string
}
