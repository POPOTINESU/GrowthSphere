package valueobject

import "errors"

type Username struct {
	value string
}

func NewUsername(username string) (Username, error) {
	if len(username) == 0 {
		return Username{}, errors.New("username can not be empty")
	}

	if len(username) > 20 {
		return Username{}, errors.New("username must be at most 20 characters")
	}

	return Username{value: username}, nil
}

func (u Username) String() string {
	return u.value
}
