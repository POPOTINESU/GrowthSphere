package valueobject

import "errors"

type AccountName struct {
	value string
}

func NewAccountName(accountName string) (AccountName, error) {
	MAX_ACCOUNT_NAME := 30

	if len(accountName) == 0 {
		return AccountName{}, errors.New("account name can not be empty")
	}

	if len(accountName) > MAX_ACCOUNT_NAME {
		return AccountName{}, errors.New("account name must be at most 30 characters")
	}

	return AccountName{value: accountName}, nil
}

func (a AccountName) String() string {
	return a.value
}
