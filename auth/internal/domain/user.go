package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Email    string
	Password *Password
	Role     Role
}
