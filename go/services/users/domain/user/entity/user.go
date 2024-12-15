package user

import (
	"GROWTHSPHERE/services/users/domain"
	valueobject "GROWTHSPHERE/services/users/domain/user/value_object"
	"errors"

	"github.com/google/uuid"
)

// User represents a user entity.
type User struct {
	domain.AggregateRoot
	FullName valueobject.FullName
}

// UserEntity defines methods for user-specific behavior.
type UserEntity interface {
	domain.AggregateRootMethod
}

// NewUser is a factory function to create a new User entity.
func NewUser(id uuid.UUID, fullName valueobject.FullName) (*User, error) {
	// Generate a new UUID if not provided
	if id == uuid.Nil {
		newID, err := uuid.NewUUID()
		if err != nil {
			return nil, errors.New("failed to generate UUID")
		}
		id = newID
	}

	// Return the constructed User
	return &User{
		AggregateRoot: domain.AggregateRoot{Id: id},
		FullName:      fullName,
	}, nil
}