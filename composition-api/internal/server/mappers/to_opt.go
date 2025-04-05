package mappers

import (
	"time"

	api "composition-api/internal/generated/http/api"
)

func ToOptString(value *string) api.OptString {
	if value == nil {
		return api.OptString{
			Set: false,
		}
	}
	return api.OptString{
		Value: *value,
		Set:   true,
	}
}

func ToOptFloat64(value *float64) api.OptFloat64 {
	if value == nil {
		return api.OptFloat64{
			Set: false,
		}
	}
	return api.OptFloat64{
		Value: *value,
		Set:   true,
	}
}

func ToOptBool(value *bool) api.OptBool {
	if value == nil {
		return api.OptBool{
			Set: false,
		}
	}
	return api.OptBool{
		Value: *value,
		Set:   true,
	}
}

func ToOptDate(value *time.Time) api.OptDate {
	if value == nil {
		return api.OptDate{
			Set: false,
		}
	}
	return api.OptDate{
		Value: *value,
		Set:   true,
	}
}
