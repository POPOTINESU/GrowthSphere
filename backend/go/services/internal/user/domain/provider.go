package domain

import "errors"

type Provider struct {
	value string
}

func NewProvider(provider string) (Provider, error) {
	switch provider {
	case "local":
		return Provider{value: provider}, nil
	default:
		return Provider{}, errors.New("not found provider")
	}
}
