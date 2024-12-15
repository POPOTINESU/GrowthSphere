package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	data := []struct {
		name     string
		password string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid password",
			password: "ValidPass123!",
			wantErr:  false,
		},
		{
			name:     "password too short",
			password: "short",
			wantErr:  true,
			errMsg:   "password must be greater than or equal to 8",
		},
		{
			name:     "password too long",
			password: string(make([]byte, 129)), // 129 characters
			wantErr:  true,
			errMsg:   "password must be less than or equal to 128",
		},
		{
			name:     "password with spaces",
			password: "  ValidPass123!  ",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  true,
			errMsg:   "password must be greater than or equal to 8",
		},
	}

	for _, testData := range data {
		t.Run(testData.name, func(t *testing.T) {
			passwordObj, err := NewPassword(testData.password)

			if testData.wantErr {
				assert.Error(t, err)
				if testData.errMsg != "" {
					assert.Contains(t, err.Error(), testData.errMsg)
				}
				assert.Empty(t, passwordObj.String())
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, passwordObj.String())
			}
		})
	}
}
