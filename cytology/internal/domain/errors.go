package domain

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrConflict            = errors.New("conflict")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)
