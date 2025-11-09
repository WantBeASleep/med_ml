package domain

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrConflict            = errors.New("conflict")
)
