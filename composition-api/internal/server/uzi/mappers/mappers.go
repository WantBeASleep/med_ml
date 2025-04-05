package mappers

type mapper[PB, D any] interface {
	Domain(pb PB) D
}

func slice[PB, D any](pbs []PB, m mapper[PB, D]) []D {
	domains := make([]D, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, m.Domain(pb))
	}
	return domains
}
