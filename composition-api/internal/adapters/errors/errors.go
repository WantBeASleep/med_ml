package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = errors.New("not found")

// HandleGRPCError обрабатывает ошибки от gRPC клиентов и конвертирует их в доменные ошибки адаптера
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
		return ErrNotFound
	default:
		return err
	}
}
