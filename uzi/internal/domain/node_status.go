package domain

import "fmt"

type NodeValidation string

const (
	// ai узел не провалидирован специалистом
	NodeValidationNull NodeValidation = "null"
	// ai узел не прошел валидацию специалистом
	NodeValidationInvalid NodeValidation = "invalid"
	// ai узел провалидирован специалистом
	NodeValidationValid NodeValidation = "valid"
)

func (s NodeValidation) String() string {
	return string(s)
}

func (s NodeValidation) Parse(status string) (NodeValidation, error) {
	switch status {
	case "null":
		return NodeValidationNull, nil
	case "invalid":
		return NodeValidationInvalid, nil
	case "valid":
		return NodeValidationValid, nil
	default:
		return "", fmt.Errorf("invalid status: %s", status)
	}
}
