package security

import (
	domain "composition-api/internal/domain/auth"

	"github.com/google/uuid"
)

type Token struct {
	Id   uuid.UUID
	Role domain.Role
}
