package domain

import "fmt"

type UziStatus string

const (
	// узи создано
	UziStatusNew UziStatus = "new"
	// узи обрабатывается
	UziStatusPending UziStatus = "pending"
	// узи обработано
	UziStatusCompleted UziStatus = "completed"
)

func (s UziStatus) String() string {
	return string(s)
}

func (s UziStatus) Parse(status string) (UziStatus, error) {
	switch status {
	case "new":
		return UziStatusNew, nil
	case "pending":
		return UziStatusPending, nil
	case "completed":
		return UziStatusCompleted, nil
	default:
		return "", fmt.Errorf("invalid status: %s", status)
	}
}
