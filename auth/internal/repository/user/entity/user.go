package entity

import (
	"database/sql"

	"github.com/WantBeASleep/med_ml_lib/gtc"

	"auth/internal/domain"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID      `db:"id"`
	Email    string         `db:"email"`
	Password sql.NullString `db:"password"`
	Role     string         `db:"role"`
}

func (e User) ToDomain() domain.User {
	var password *domain.Password
	if e.Password.Valid {
		passwordParsed, _ := domain.Password{}.Parse(e.Password.String)
		password = &passwordParsed
	}

	return domain.User{
		Id:       e.Id,
		Email:    e.Email,
		Password: password,
		Role:     domain.Role(e.Role),
	}
}

func (User) FromDomain(d domain.User) User {
	var password *string
	if d.Password != nil {
		passwordStr := d.Password.String()
		password = &passwordStr
	}

	return User{
		Id:       d.Id,
		Email:    d.Email,
		Password: gtc.String.PointerToSql(password),
		Role:     d.Role.String(),
	}
}
