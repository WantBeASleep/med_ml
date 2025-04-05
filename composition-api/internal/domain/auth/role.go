package domain

import "fmt"

type Role string

const (
	RoleDoctor  Role = "doctor"
	RolePatient Role = "patient"
)

func (Role) Parse(str string) (Role, error) {
	switch str {
	case string(RoleDoctor):
		return RoleDoctor, nil
	case string(RolePatient):
		return RolePatient, nil
	}
	return "", fmt.Errorf("invalid role: %s", str)
}

func (r Role) String() string {
	return string(r)
}
