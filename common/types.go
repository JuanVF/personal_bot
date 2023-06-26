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
package common

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type EnvironmentConfig struct {
	Development Configuration `yaml:"development"`
	Container   Configuration `yaml:"container"`
	Test        Configuration `yaml:"test"`
}

type Configuration struct {
	Google        GoogleConf        `yaml:"google"`
	Bot           BotConf           `yaml:"bot"`
	PersonalBotDB PersonalBotDBConf `yaml:"personal_bot_db"`
	OpenAI        OpenAIConfig      `yaml:"open_ai"`
	AWS           AWS               `yaml:"aws"`
}

type AWS struct {
	SES SES `yaml:"ses"`
}

type SES struct {
	Sender         string `yaml:"sender"`
	BudgetTemplate string `yaml:"budget_template"`
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
	mock       *sqlmock.Sqlmock
}

type Response struct {
	Status int
	Body   any
}

type ErrorResponse struct {
	Message string
}
