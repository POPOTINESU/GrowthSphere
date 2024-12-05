package valueobject

import (
	valueobject "GROWTHSPHERE/services/users/domain/valueObject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountName(t *testing.T) {
	tests := []struct {
		name         string
		accountName  string
		want         string
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "Happy Path: correct account name",
			accountName:  "test_account_name",
			want:         "test_account_name",
			expectErr:    false,
			expectErrMsg: "",
		},
		{
			name:         "Negative: account name is empty",
			accountName:  "",
			want:         "",
			expectErr:    true,
			expectErrMsg: "account name cannot be empty",
		},
		{
			name:         "Negative: username exceeds max length",
			accountName:  "this_account_name_is_way_too_long_therefore_account_name_is_not_valid",
			expectErr:    true,
			expectErrMsg: "account name must be at most 30 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := valueobject.NewAccountName(tt.accountName)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectErrMsg, err.Error(), "error message should match")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String(), "username should match")
			}
		})
	}
}
