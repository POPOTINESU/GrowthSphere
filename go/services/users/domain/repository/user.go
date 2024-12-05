package repository

import (
	"GROWTHSPHERE/services/users/domain/entity"
	"context"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entity.User) (uuid.UUID, error)
}
