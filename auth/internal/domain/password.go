package domain

import "fmt"

type Password struct {
	Hash string
	Salt string
}

func (Password) Parse(str string) (Password, error) {
	if len(str) < 64 {
		return Password{}, fmt.Errorf("password is too short")
	}

	hash := str[:64]
	salt := str[64:]

	return Password{Hash: hash, Salt: salt}, nil
}

func (p Password) String() string {
	return fmt.Sprintf("%s%s", p.Hash, p.Salt)
}
