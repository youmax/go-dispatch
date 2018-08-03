package responses

import (
	"errors"
	"net/http"
)

var (
	// ErrInvalidParam either means the given route parameter was wrong, like a non uint, or too long
	ErrInvalidParam  = &RequestError{ErrorString: "Bad Request", ErrorCode: http.StatusBadRequest}
	ErrInternalError = &RequestError{ErrorString: "Internal Error", ErrorCode: http.StatusInternalServerError}
	ErrNotFound      = &RequestError{ErrorString: "Request Not Found", ErrorCode: http.StatusNotFound}
	// ErrUnauthorized means the user could not be validated and any JWT tokens on client side should be removed
	ErrUnauthorized = &RequestError{ErrorString: "Unauthorized", ErrorCode: http.StatusUnauthorized}
	// ErrForbidden is either anon accessing a route that requires auth, or an authed user without the correct permissions
	ErrForbidden = &RequestError{ErrorString: "Forbidden", ErrorCode: http.StatusForbidden}

	ErrTokenInvalid = errors.New("Invalid token")
	ErrUserNotValid = errors.New("User is not valid")
	ErrCsrfNotValid = errors.New("CSRF token is not valid")
)

// RequestError holds the message string and http code
type RequestError struct {
	ErrorString string
	ErrorCode   int
}

// Code returns the http error code
func (err *RequestError) Code() int {
	return err.ErrorCode
}

func (err *RequestError) Error() string {
	return err.ErrorString
}

// NewError returns the code and message for Gins JSON helpers
func NewError(e interface{}) (int, map[string]interface{}) {
	switch e.(type) {
	case string:
		return 400, map[string]interface{}{"message": e, "status_code": 400}
	case *RequestError:
		e, _ := e.(*RequestError)
		return e.Code(), map[string]interface{}{"message": e.Error(), "status_code": e.Code()}
	default:
		return 400, map[string]interface{}{"message": "error", "status_code": 400}
	}

}
