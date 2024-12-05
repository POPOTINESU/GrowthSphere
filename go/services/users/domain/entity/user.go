package entity

import (
	valueobject "GROWTHSPHERE/services/users/domain/valueObject"
	"time"

	"github.com/google/uuid"
)

type User struct {
	userID      uuid.UUID
	username    valueobject.Username
	accountName valueobject.AccountName
	updatedAt   time.Time
}

// NewUser creates a new User entity with the given username and account name.
// It validates the inputs and generates a unique identifier for the user.
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

// UserID returns the unique identifier of the user.
func (u *User) UserID() uuid.UUID {
	return u.userID
}

// Username returns the username of the user.
func (u *User) Username() valueobject.Username {
	return u.username
}

// AccountName returns the account name associated with the user.
func (u *User) AccountName() valueobject.AccountName {
	return u.accountName
}

// UpdatedAt returns the timestamp of the last update for the user.
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
