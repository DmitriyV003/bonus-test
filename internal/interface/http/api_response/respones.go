package api_response

type ErrorApiResponse struct {
	Message string `json:"message"`
}

func (e *ErrorApiResponse) Error() string {
	return e.Message
}

func CreateErrorApiResponse(message string) ErrorApiResponse {
	return ErrorApiResponse{Message: message}
}
