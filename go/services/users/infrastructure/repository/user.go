package userRepository

import (
	"GROWTHSPHERE/ent"
	"GROWTHSPHERE/services/users/domain/entity"
	"GROWTHSPHERE/services/users/domain/repository"
	"context"

	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.IUserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	uuid, _ := uuid.NewV7()
	return uuid, nil
}
