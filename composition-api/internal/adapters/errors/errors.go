package errors

import (
	"fmt"

	"composition-api/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGRPCError обрабатывает ошибки от gRPC клиентов и конвертирует их в доменные ошибки
func HandleGRPCError(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return fmt.Errorf("unknown error: %w", err)
	}

	switch st.Code() {
	case codes.NotFound:
		return domain.ErrNotFound
	case codes.InvalidArgument:
		return domain.ErrBadRequest
	case codes.Unauthenticated:
		return domain.ErrUnauthorized
	case codes.AlreadyExists:
		return domain.ErrConflict
	case codes.FailedPrecondition:
		return domain.ErrUnprocessableEntity
	default:
		return err
	}
}
