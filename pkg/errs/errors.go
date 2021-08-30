package errs

import "errors"

var DbError = errors.New("unexpected database error")
var ContentType404 = errors.New("content type not found")
var Column404 = errors.New("column not found")
var Content404 = errors.New("content not found")
