package errs

import "errors"

var ErrDb = errors.New("unexpected database error")
var ErrContentTypeNotFound = errors.New("content type not found")
var ErrColumnNotFound = errors.New("column not found")
var ErrContentNotFound = errors.New("content not found")
var ErrEmailMissing = errors.New("email is missing")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidEmail = errors.New("invalid email")
var ErrServerErr = errors.New("internal server error")
var ErrTokenErr = errors.New("can't generate token")
var ErrContentWithStatusOk = errors.New("there is no content found")
var ErrContentParams = errors.New("parameter value is not valid")
var ErrUnexpected = errors.New("unexpected error")
var ErrInvalidToken = errors.New("invalid token")

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}


