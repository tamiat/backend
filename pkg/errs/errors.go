package errs

import "errors"

var (
	ErrDb = errors.New("unexpected database error")
	ErrContentTypeNotFound = errors.New("content type not found")
	ErrColumnNotFound = errors.New("column not found")
	ErrContentNotFound = errors.New("content not found")
)

