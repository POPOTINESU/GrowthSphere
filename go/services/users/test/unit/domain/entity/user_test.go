package entity

import (
	"GROWTHSPHERE/services/users/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
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
			name:         "Negative: invalid username",
			username:     "",
			accountName:  "valid_account",
			expectErr:    true,
			expectErrMsg: "username cannot be empty",
		},
		{
			name:         "Negative: invalid account name",
			username:     "valid_user",
			accountName:  "",
			expectErr:    true,
			expectErrMsg: "account name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entity.NewUser(tt.username, tt.accountName)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectErrMsg, err.Error(), "error message should match")
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user, "user should not be nil")
				assert.NotNil(t, user.UserId(), "userID should not be nil")
				assert.Equal(t, tt.username, user.Username().String(), "username should match")
				assert.Equal(t, tt.accountName, user.AccountName().String(), "accountName should match")
				assert.WithinDuration(t, time.Now(), user.UpdatedAt(), time.Second, "updatedAt should be set to current time")
			}
		})
	}
}
