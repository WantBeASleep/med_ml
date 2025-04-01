package user

type registerOptions struct {
	password *string
}

type registerOption func(*registerOptions)

var (
	WithPassword = func(password string) registerOption {
		return func(o *registerOptions) {
			o.password = &password
		}
	}
)
