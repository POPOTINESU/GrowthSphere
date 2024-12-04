package repository

import (
	"GROWTHSPHERE/services/users/internal/domain/entity"
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
}
