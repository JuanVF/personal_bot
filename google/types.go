package google

type OAuth2Response struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type GmailThreads struct {
	Messages           []GmailThreadIdentifier `json:"messages"`
	NextPageToken      string                  `json:"nextPageToken"`
	ResultSizeEstimate int                     `json:"resultSizeEstimate"`
}

type GmailThreadIdentifier struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
}
