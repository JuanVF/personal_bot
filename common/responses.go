package common

// Constructor for getting an error response
func GetErrorResponse(message string, status int) *Response {
	return &Response{
		Status: status,
		Body: &ErrorResponse{
			Message: message,
		},
	}
}
