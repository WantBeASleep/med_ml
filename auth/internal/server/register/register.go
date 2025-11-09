package register

import (
	"context"
	"errors"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
	"auth/internal/server/mappers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RegisterUser(ctx context.Context, in *pb.RegisterUserIn) (*pb.RegisterUserOut, error) {
	id, err := h.services.UserService.RegisterUser(
		ctx,
		in.Email,
		in.Password,
		mappers.RoleReversedMap[in.Role],
	)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return nil, status.Errorf(codes.InvalidArgument, "Неверный формат запроса")
		case errors.Is(err, domain.ErrConflict):
			return nil, status.Errorf(codes.AlreadyExists, "Пользователь с таким email уже существует")
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return nil, status.Errorf(codes.FailedPrecondition, "Ошибка валидации данных")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.RegisterUserOut{Id: id.String()}, nil
}
