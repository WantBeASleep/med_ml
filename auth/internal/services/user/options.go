package user

import "auth/internal/domain"

type registerOptions struct {
	password *string
	role     *domain.Role
}

type registerOption func(*registerOptions)

var (
	WithPassword = func(password string) registerOption {
		return func(o *registerOptions) {
			o.password = &password
		}
	}
	WithRole = func(role domain.Role) registerOption {
		return func(o *registerOptions) {
			o.role = &role
		}
	}
)
