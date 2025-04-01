package utils

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func MustParseFromClaims(t *testing.T, key string, claims jwt.MapClaims) string {
	t.Helper()

	valueAny, ok := claims[key]
	require.True(t, ok)

	value, ok := valueAny.(string)
	require.True(t, ok)

	return value
}
