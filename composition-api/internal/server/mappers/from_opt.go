package mappers

import (
	api "composition-api/internal/generated/http/api"
)

func FromOptString(opt api.OptString) *string {
	if !opt.Set {
		return nil
	}
	return &opt.Value
}

func FromOptFloat64(opt api.OptFloat64) *float64 {
	if !opt.Set {
		return nil
	}
	return &opt.Value
}

func FromOptBool(opt api.OptBool) *bool {
	if !opt.Set {
		return nil
	}
	return &opt.Value
}
