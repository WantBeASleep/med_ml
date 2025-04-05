package domain

import "fmt"

type UziProjection string

const (
	UziProjectionLong  UziProjection = "long"
	UziProjectionCross UziProjection = "cross"
)

func (s UziProjection) String() string {
	return string(s)
}

func (s UziProjection) Parse(projection string) (UziProjection, error) {
	switch projection {
	case "long":
		return UziProjectionLong, nil
	case "cross":
		return UziProjectionCross, nil
	default:
		return "", fmt.Errorf("invalid projection: %s", projection)
	}
}
