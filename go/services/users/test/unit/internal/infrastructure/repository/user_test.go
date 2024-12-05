package repository_test

import (
	entUser "GROWTHSPHERE/ent/user"
	"GROWTHSPHERE/pkg/db"
	"GROWTHSPHERE/services/users/domain/entity"
	userRepository "GROWTHSPHERE/services/users/infrastructure/repository"
	"context"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := db.CreateTestClient()
	err := client.Schema.Create(context.Background())
	assert.NoError(t, err, "failed to run migrations")

	userRepo := userRepository.NewUserRepository(client)

	tests := []struct {
		name         string
		username     string
		accountName  string
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "Happy Path: valid inputs",
			username:     "valid_user",
			accountName:  "valid_account",
			expectErr:    false,
			expectErrMsg: "",
		},
		{
			name:         "Negative: missing username",
			username:     "",
			accountName:  "valid_account",
			expectErr:    true,
			expectErrMsg: "username can not be empty",
		},
		{
			name:         "Negative: missing account name",
			username:     "valid_user",
			accountName:  "",
			expectErr:    true,
			expectErrMsg: "account name can not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entity.NewUser(tt.username, tt.accountName)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectErrMsg, err.Error(), "error message should match")
				return
			}

			assert.NoError(t, err, "failed to create user entity")
			assert.NotNil(t, user, "user should not be nil")

			createdUserID, err := userRepo.Create(context.Background(), user)

			assert.NoError(t, err, "unexpected error during user creation")
			assert.NotNil(t, createdUserID, "created user ID should not be nil")

			dbUser, err := client.User.Query().Where(entUser.IDEQ(createdUserID)).Only(context.Background())
			assert.NoError(t, err, "failed to query user from database")
			assert.Equal(t, user.Username().String(), dbUser.Username, "username in database should match")
			assert.Equal(t, user.AccountName().String(), dbUser.AccountName, "account name in database should match")
			assert.WithinDuration(t, time.Now(), dbUser.UpdatedAt, time.Second, "updatedAt in database should be close to current time")
		})
	}
}
