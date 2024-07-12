package response

// Response is a model for API responses
// swagger:model Response
type Response struct {
	// The response message
	// example: "Request was successful"
	Message string `json:"message"`
	// The response data
	// example: {"key": "value"}
	Data interface{} `json:"data"`
	// The error information, if any
	// example: "Error description"
	Error interface{} `json:"error"`
}

func ClientResponse(message string, data interface{}, err interface{}) Response {

	return Response{
		Message: message,
		Data:    data,
		Error:   err,
	}

}
