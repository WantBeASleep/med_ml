package entity

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrValidation = errors.New("validation error")
)

// WrapDBError оборачивает ошибку БД в доменную ошибку repository.
// Использует errors.As для проверки конкретных типов ошибок БД согласно Uber Go Style Guide.
func WrapDBError(err error) error {
	if err == nil {
		return nil
	}

	// Проверяем ошибки PostgreSQL через errors.As
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		// Коды ошибок PostgreSQL для constraint violations и check constraints
		switch pqErr.Code {
		case pq.ErrorCode("23514"), // check_violation
			pq.ErrorCode("23505"), // unique_violation
			pq.ErrorCode("23503"), // foreign_key_violation
			pq.ErrorCode("23502"): // not_null_violation
			return fmt.Errorf("%w: %w", ErrValidation, err)
		}
	}

	// Возвращаем оригинальную ошибку, если не удалось определить тип
	return err
}
