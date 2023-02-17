package common

type Configuration struct {
	Gmail GmailConf `yaml:"gmail"`
}

type GmailConf struct {
	BaseURL     string `yaml:"base_url"`
	ClientId    string `yaml:"client_id"`
	Secret      string `yaml:"secret"`
	RedirectURI string `yaml:"redirect_uri"`
}

type Logger struct {
	headline string
}
