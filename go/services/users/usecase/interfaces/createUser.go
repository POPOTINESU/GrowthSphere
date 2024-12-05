package interfaces

import (
	"context"

	"github.com/google/uuid"
)

// ICreateUserUsecase defines the interface for creating a user.
type ICreateUserUsecase interface {
	Execute(ctx context.Context, username string, accountName string) (uuid.UUID, error)
}
