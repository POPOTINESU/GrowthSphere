package domain

import (
	"errors"
)

type Password struct {
	value string
}

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, errors.New("password must be at least 8 characters")
	}

	return Password{value: password}, nil
}
