package userRepository

import (
	"GROWTHSPHERE/ent"
	"GROWTHSPHERE/services/users/domain/entity"
	"GROWTHSPHERE/services/users/domain/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.IUserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	createdUser, err := r.client.User.
		Create().
		SetID(user.UserID()).
		SetAccountName(user.AccountName().String()).
		SetUsername(user.Username().String()).
		SetUpdatedAt(user.UpdatedAt()).Save(ctx)

	if err != nil {
		fmt.Println("cannot create user:", err)
		return uuid.Nil, errors.New("cannot create user")
	}

	return createdUser.ID, nil
}
