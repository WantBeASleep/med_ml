package mappers

import (
	api "gateway/internal/generated/http/api"
)

func Contor(contor []byte) (api.Contor, error) {
	out := &api.Contor{}

	err := out.UnmarshalJSON(contor)
	if err != nil {
		return api.Contor{}, err
	}

	return *out, nil
}
