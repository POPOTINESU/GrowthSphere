package infrastructure

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/argon2"
)

const (
	// Default Argon2 settings
	DefaultMemory      = 64 * 1024
	DefaultIterations  = 3
	DefaultParallelism = 2
	DefaultKeyLength   = 32
	DefaultSaltLength  = 16
)

type Argon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	KeyLength   uint32
	SaltLength  uint32
}

// LoadArgon2Config loads Argon2 settings from environment variables or defaults.
func LoadArgon2Config() (Argon2Config, error) {
	parseUint := func(envVar string, defaultValue uint64, maxValue uint64) (uint64, error) {
		if value := os.Getenv(envVar); value != "" {
			parsed, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return 0, err
			}
			if parsed > maxValue {
				return 0, errors.New(envVar + " exceeds maximum allowed value")
			}
			return parsed, nil
		}
		return defaultValue, nil
	}

	// Add explicit range checks before casting
	checkAndCastToUint32 := func(value uint64) (uint32, error) {
		if value > uint64(^uint32(0)) {
			return 0, errors.New("value exceeds uint32 max")
		}
		return uint32(value), nil
	}

	checkAndCastToUint8 := func(value uint64) (uint8, error) {
		if value > uint64(^uint8(0)) {
			return 0, errors.New("value exceeds uint8 max")
		}
		return uint8(value), nil
	}

	memory, err := parseUint("ARGON2_MEMORY", DefaultMemory, uint64(^uint32(0)))
	if err != nil {
		return Argon2Config{}, err
	}
	memoryUint32, err := checkAndCastToUint32(memory)
	if err != nil {
		return Argon2Config{}, err
	}

	iterations, err := parseUint("ARGON2_ITERATIONS", DefaultIterations, uint64(^uint32(0)))
	if err != nil {
		return Argon2Config{}, err
	}
	iterationsUint32, err := checkAndCastToUint32(iterations)
	if err != nil {
		return Argon2Config{}, err
	}

	parallelism, err := parseUint("ARGON2_PARALLELISM", DefaultParallelism, uint64(^uint8(0)))
	if err != nil {
		return Argon2Config{}, err
	}
	parallelismUint8, err := checkAndCastToUint8(parallelism)
	if err != nil {
		return Argon2Config{}, err
	}

	keyLength, err := parseUint("ARGON2_KEY_LENGTH", DefaultKeyLength, uint64(^uint32(0)))
	if err != nil {
		return Argon2Config{}, err
	}
	keyLengthUint32, err := checkAndCastToUint32(keyLength)
	if err != nil {
		return Argon2Config{}, err
	}

	return Argon2Config{
		Memory:      memoryUint32,
		Iterations:  iterationsUint32,
		Parallelism: parallelismUint8,
		KeyLength:   keyLengthUint32,
		SaltLength:  DefaultSaltLength,
	}, nil
}

// HashPassword hashes the password using Argon2id and the given salt.
func HashPassword(password string, salt []byte) (string, error) {
	config, err := LoadArgon2Config()
	if err != nil {
		return "", err
	}

	// Generate Argon2 hash
	hash := argon2.IDKey([]byte(password), salt, config.Iterations, config.Memory, config.Parallelism, config.KeyLength)

	// Combine salt and hash, and encode as base64
	return base64.StdEncoding.EncodeToString(append(salt, hash...)), nil
}

// GenerateSalt generates a cryptographically secure random salt.
func GenerateSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt) // Use crypto/rand for secure random bytes
	if err != nil {
		return nil, err
	}
	return salt, nil
}
