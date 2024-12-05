package valueobject

import (
	valueobject "GROWTHSPHERE/services/users/domain/valueObject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsername(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		want         string
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "Happy Path: correct username",
			username:     "test_user",
			want:         "test_user",
			expectErr:    false,
			expectErrMsg: "",
		},
		{
			name:         "Negative: user name is empty",
			username:     "",
			want:         "",
			expectErr:    true,
			expectErrMsg: "username cannot be empty",
		},
		{
			name:         "Negative: username exceeds max length",
			username:     "this_username_is_way_too_long",
			want:         "",
			expectErr:    true,
			expectErrMsg: "username must be at most 20 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := valueobject.NewUsername(tt.username)

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
