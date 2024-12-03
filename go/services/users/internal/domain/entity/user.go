package entity

import (
	valueobject "GROWTHSPHERE/services/users/internal/domain/valueObject"
	"time"

	"github.com/google/uuid"
)

type User struct {
	userID      uuid.UUID
	username    valueobject.Username
	accountName valueobject.AccountName
	updatedAt   time.Time
}

func NewUser(username string, accountName string) (*User, error) {
	cleanUsername, err := valueobject.NewUsername(username)
	if err != nil {
		return &User{}, err
	}

	cleanAccountName, err := valueobject.NewAccountName(accountName)
	if err != nil {
		return &User{}, err
	}

	newUuid, err := uuid.NewV7()
	if err != nil {
		return &User{}, err
	}

	return &User{
		userID:      newUuid,
		username:    cleanUsername,
		accountName: cleanAccountName,
		updatedAt:   time.Now(),
	}, nil
}

func (u *User) UserId() uuid.UUID {
	return u.userID
}

func (u *User) Username() valueobject.Username {
	return u.username
}

func (u *User) AccountName() valueobject.AccountName {
	return u.accountName
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
