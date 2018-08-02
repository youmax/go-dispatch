package responses

// ErrorMessage returns the code and message for Gins JSON helpers
func NewResponse(statusCode int, message map[string]interface{}) (int, map[string]interface{}) {
	message["status_code"] = statusCode
	return statusCode, message
}
