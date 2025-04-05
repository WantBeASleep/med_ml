package token

import "time"

type generateTokenOptions struct {
	expirationTime *time.Time
}

type generateTokenOption func(*generateTokenOptions)

var WithExpirationTime = func(expirationTime time.Time) generateTokenOption {
	return func(o *generateTokenOptions) {
		o.expirationTime = &expirationTime
	}
}
