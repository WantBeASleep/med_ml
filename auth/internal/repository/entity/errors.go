package entity

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("not found")
)

// DBConflictError представляет ошибку конфликта в БД (unique_violation).
type DBConflictError struct {
	Err error
}

func (e *DBConflictError) Error() string {
	return "conflict: " + e.Err.Error()
}

func (e *DBConflictError) Unwrap() error {
	return e.Err
}

// DBValidationError представляет ошибку валидации в БД (check_violation, foreign_key_violation, not_null_violation).
type DBValidationError struct {
	Err error
}

func (e *DBValidationError) Error() string {
	return "validation error: " + e.Err.Error()
}

func (e *DBValidationError) Unwrap() error {
	return e.Err
}

// WrapDBError оборачивает ошибку БД в типизированную ошибку repository.
// Использует errors.As для проверки конкретных типов ошибок БД согласно Uber Go Style Guide.
// Возвращает типизированные ошибки, которые можно проверить через errors.As().
func WrapDBError(err error) error {
	if err == nil {
		return nil
	}

	// Проверяем ошибки PostgreSQL через errors.As
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		// Коды ошибок PostgreSQL для constraint violations и check constraints
		switch pqErr.Code {
		case pq.ErrorCode("23505"): // unique_violation
			return &DBConflictError{Err: err}
		case pq.ErrorCode("23514"), // check_violation
			pq.ErrorCode("23503"), // foreign_key_violation
			pq.ErrorCode("23502"): // not_null_violation
			return &DBValidationError{Err: err}
		}
	}

	// Возвращаем оригинальную ошибку, если не удалось определить тип
	return err
}
