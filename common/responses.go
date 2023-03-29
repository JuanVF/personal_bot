package common

import "net/http"

// Constructor for getting an error response
func GetErrorResponse(message string, status int) *Response {
	return &Response{
		Status: status,
		Body: &ErrorResponse{
			Message: message,
		},
	}
}

// Constructor for getting an 200 OK response with any body
func GetSuccessResponse(body interface{}) *Response {
	return &Response{
		Status: http.StatusOK,
		Body:   body,
	}
}
