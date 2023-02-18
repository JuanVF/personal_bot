package common

import "database/sql"

type Configuration struct {
	Google        GoogleConf        `yaml:"google"`
	Bot           BotConf           `yaml:"bot"`
	PersonalBotDB PersonalBotDBConf `yaml:"personal_bot_db"`
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
