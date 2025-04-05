package auth

import (
	"context"

	domain "composition-api/internal/domain/auth"
	pb "composition-api/internal/generated/grpc/clients/auth"

	"github.com/google/uuid"
)

type Adapter interface {
	Login(ctx context.Context, email, password string) (domain.Token, domain.Token, error)
	Refresh(ctx context.Context, refreshToken domain.Token) (domain.Token, domain.Token, error)
	RegisterUser(ctx context.Context, email, password string, role domain.Role) (uuid.UUID, error)
	CreateUnRegisteredUser(ctx context.Context, email string) (uuid.UUID, error)
}

type adapter struct {
	client pb.AuthSrvClient
}

func NewAdapter(client pb.AuthSrvClient) Adapter {
	return &adapter{client: client}
}
