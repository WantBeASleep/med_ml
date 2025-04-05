package mappers

type mapper[D, Api any] interface {
	Api(d D) Api
}

func slice[D, Api any](ds []D, m mapper[D, Api]) []Api {
	apis := make([]Api, 0, len(ds))
	for _, d := range ds {
		apis = append(apis, m.Api(d))
	}
	return apis
}
