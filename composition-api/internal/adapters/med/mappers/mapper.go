package mappers

import "github.com/AlekSi/pointer"

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

func PointerFromMap[M comparable, N any](m map[M]N, key *M) *N {
	if key == nil {
		return nil
	}
	return pointer.To(m[*key])
}
