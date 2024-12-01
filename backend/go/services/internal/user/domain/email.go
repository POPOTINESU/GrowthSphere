package domain

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(email string) (Email, error) {
	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("email is not formatted correctly")
	}

	return Email{value: email}, nil
}
