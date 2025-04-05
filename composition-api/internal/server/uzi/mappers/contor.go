package mappers

import (
	api "composition-api/internal/generated/http/api"
)

func Contor(contor []byte) (api.Contor, error) {
	out := &api.Contor{}

	err := out.UnmarshalJSON(contor)
	if err != nil {
		return api.Contor{}, err
	}

	return *out, nil
}
