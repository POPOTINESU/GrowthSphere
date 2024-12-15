package valueobject

import (
	"GROWTHSPHERE/pkg/validation"
	"GROWTHSPHERE/services/auth/infrastructure"
	"strings"
)

type Password struct {
	hashedPassword string
}

const (
	PasswordMinLength = 8
	PasswordMaxLength = 128

	// Default Argon2 settings
	DefaultMemory      = 64 * 1024 // 64MB
	DefaultIterations  = 3
	DefaultParallelism = 2
	DefaultKeyLength   = 32
	DefaultSaltLength  = 16
)

// NewPassword creates a new hashed Password Value Object.
func NewPassword(password string) (Password, error) {
	// Trim spaces from the password
	password = strings.TrimSpace(password)

	if err := validation.IsInRange(len(password), PasswordMinLength, PasswordMaxLength, "password"); err != nil {
		return Password{}, err
	}

	// Generate secure random salt
	salt, err := infrastructure.GenerateSalt(DefaultSaltLength)
	if err != nil {
		return Password{}, err
	}

	// Hash the password using Argon2
	hashedPassword, err := infrastructure.HashPassword(password, salt)
	if err != nil {
		return Password{}, err
	}

	return Password{hashedPassword: hashedPassword}, nil
}

// GetHashedPassword returns the hashed password as a string.
func (p Password) String() string {
	return p.hashedPassword
}
