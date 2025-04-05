package entity

import (
	"database/sql"

	"github.com/WantBeASleep/med_ml_lib/gtc"

	"auth/internal/domain"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID      `db:"id"`
	Email        string         `db:"email"`
	Password     sql.NullString `db:"password"`
	RefreshToken sql.NullString `db:"refresh_token"`
	Role         string         `db:"role"`
}

func (e User) ToDomain() domain.User {
	var password *domain.Password
	if e.Password.Valid {
		passwordParsed, _ := domain.Password{}.Parse(e.Password.String)
		password = &passwordParsed
	}

	var refreshToken *domain.Token
	if e.RefreshToken.Valid {
		refreshTokenParsed := domain.Token(e.RefreshToken.String)
		refreshToken = &refreshTokenParsed
	}

	return domain.User{
		Id:           e.Id,
		Email:        e.Email,
		Password:     password,
		RefreshToken: refreshToken,
		Role:         domain.Role(e.Role),
	}
}

func (User) FromDomain(d domain.User) User {
	var password *string
	if d.Password != nil {
		passwordStr := d.Password.String()
		password = &passwordStr
	}

	var refreshToken *string
	if d.RefreshToken != nil {
		refreshTokenStr := d.RefreshToken.String()
		refreshToken = &refreshTokenStr
	}

	return User{
		Id:           d.Id,
		Email:        d.Email,
		Password:     gtc.String.PointerToSql(password),
		RefreshToken: gtc.String.PointerToSql(refreshToken),
		Role:         d.Role.String(),
	}
}
