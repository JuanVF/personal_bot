package google

type OAuth2Response struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
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

type GmailMessage struct {
	Id           string              `json:"id"`
	ThreadId     string              `json:"threadId"`
	LabelIds     []string            `json:"labelIds"`
	Snippet      string              `json:"snippet"`
	Payload      GmailMessagePayload `json:"payload"`
	SizeEstimate int                 `json:"sizeEstimate"`
	HistoryId    string              `json:"historyId"`
	InternalDate string              `json:"internalDate"`
}

type GmailMessagePayload struct {
	PartId   string                `json:"partId"`
	MimeType string                `json:"mimeType"`
	Filename string                `json:"filename"`
	Headers  []GmailMessageHeader  `json:"headers"`
	Body     GmailMessageBody      `json:"body"`
	Parts    []GmailMessagePayload `json:"parts"`
}

type GmailMessageHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GmailMessageBody struct {
	Size int    `json:"size"`
	Data string `json:"data"`
}

type VerifyTokenResponse struct {
	Azp        string `json:"azp"`
	Aud        string `json:"aud"`
	Scope      string `json:"scope"`
	Exp        string `json:"exp"`
	ExpiresIn  string `json:"expires_in"`
	AccessType string `json:"access_type"`
}

type IDTokenInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
}
