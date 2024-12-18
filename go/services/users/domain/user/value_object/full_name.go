package valueobject

import (
	"GROWTHSPHERE/pkg/validation"
	"strings"
)

type FullName struct {
	firstName string
	lastName  string
}

const (
	// first name and last name range
	NameMinLength = 1
	NameMaxLength = 40
)

// NewFullName is a factory function to create new FullName value object.
func NewFullName(firstName string, lastName string) (FullName, error) {

	// trim spaces from first name and last name
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	// first name validate
	if err := validation.IsInRange(len(firstName), NameMinLength, NameMaxLength, "first name"); err != nil {
		return FullName{}, err
	}

	// last name validate
	if err := validation.IsInRange(len(lastName), NameMinLength, NameMaxLength, "last name"); err != nil {
		return FullName{}, err
	}
	return FullName{firstName: firstName, lastName: lastName}, nil
}

func (fn FullName) String() string {
	return fn.firstName + " " + fn.lastName
}
