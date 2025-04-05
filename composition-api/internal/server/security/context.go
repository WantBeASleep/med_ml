package security

import (
	"context"
)

type key struct{}

var tokenKey = key{}

func ParseToken(ctx context.Context) (Token, error) {
	token, ok := ctx.Value(tokenKey).(Token)
	if !ok {
		return Token{}, ErrUnauthorized
	}

	return token, nil
}
