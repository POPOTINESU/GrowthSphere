package interfaces

import (
	"context"

	"github.com/google/uuid"
)

// CreateUserUsecase defines the interface for creating a user.
type CreateUserUsecase interface {
	Execute(ctx context.Context, username string, accountName string) (uuid.UUID, error)
}
