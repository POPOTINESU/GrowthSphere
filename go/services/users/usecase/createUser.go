package usecase

import (
	"GROWTHSPHERE/services/users/domain/entity"
	"GROWTHSPHERE/services/users/domain/repository"
	"context"

	"github.com/google/uuid"
)

type CreateUserUsecase struct {
	userRepo repository.IUserRepository
}

// NewCreateUserUsecase creates a new instance of ICreateUserUsecase
func NewCreateUserUsecase(userRepo repository.IUserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo: userRepo,
	}
}

// Execute creates a new user and returns its UUID
func (c *CreateUserUsecase) Execute(ctx context.Context, username string, accountName string) (uuid.UUID, error) {
	userDomain, err := entity.NewUser(username, accountName)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := c.userRepo.Create(ctx, userDomain)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
