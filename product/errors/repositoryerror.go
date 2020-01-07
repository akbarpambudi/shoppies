package errors

import "errors"

var (
	ProductNotFoundErr      = errors.New("product not found")
	ProductAlreadyExistsErr = errors.New("product already exists")
)
