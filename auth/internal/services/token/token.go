package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var ErrExpiredToken = errors.New("token expired")

func (s *service) generateToken(data map[string]any, opts ...generateTokenOption) (string, error) {
	options := &generateTokenOptions{}
	for _, opt := range opts {
		opt(options)
	}

	claims := jwt.MapClaims(data)

	if options.expirationTime != nil {
		claims["exp"] = options.expirationTime.Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *service) parseClaims(token string) (map[string]any, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return s.publicKey, nil })
	if err != nil || !parsed.Valid {
		return nil, ErrExpiredToken
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid jwt claims")
	}

	return claims, nil
}
