package errs

import "errors"

var DbError = errors.New("unexpected database error")
var ContentType404 = errors.New("content type not found")
var Column404 = errors.New("column not found")
var Content404 = errors.New("content not found")
var EmailMissing = errors.New("email is missing")
var InvalidPassword = errors.New("invalid password")
var InvalidEmail = errors.New("invalid email")
var ServerErr = errors.New("internal server error")
var TokenErr = errors.New("can't generate token")
var Content200 = errors.New("there is no content found")
var ContentParams = errors.New("parameter value is not valid")
var UnexpectedError = errors.New("unexpected error")
var InvalidToken = errors.New("invalid token")

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
