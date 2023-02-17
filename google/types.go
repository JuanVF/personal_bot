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
	Size int `json:"size"`
}
