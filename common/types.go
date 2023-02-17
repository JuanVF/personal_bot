package common

type Configuration struct {
	Google GoogleConf `yaml:"google"`
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
