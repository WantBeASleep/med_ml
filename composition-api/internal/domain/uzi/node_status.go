package domain

import "fmt"

type NodeValidation string

const (
	// узел не оценен специалистом
	NodeValidationNull NodeValidation = "null"
	// узел не оценен специалистом
	NodeValidationInvalid NodeValidation = "invalid"
	// узел оценен специалистом
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
