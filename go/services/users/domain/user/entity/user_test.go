package user

import (
	valueobject "GROWTHSPHERE/services/users/domain/user/value_object"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(test *testing.T) {
	validFullName, _ := valueobject.NewFullName("User", "Name")

	data := []struct {
		name      string
		id        uuid.UUID
		fullName  valueobject.FullName
		expectErr bool
		errMsg    string
	}{
		{
			name:      "valid user with provided UUID",
			id:        uuid.New(),
			fullName:  validFullName,
			expectErr: false,
			errMsg:    "",
		},
		{
			name:      "valid user with generated UUID",
			id:        uuid.Nil,
			fullName:  validFullName,
			expectErr: false,
			errMsg:    "",
		},
	}

	for _, testData := range data {
		test.Run(testData.name, func(test *testing.T) {
			userObj, err := NewUser(testData.id, testData.fullName)

			if testData.expectErr {
				assert.Error(test, err)
				assert.Contains(test, err.Error(), testData.errMsg, "error message should match")
				assert.Nil(test, userObj)
			} else {
				assert.NoError(test, err)
				assert.NotNil(test, userObj)
				assert.Equal(test, testData.fullName, userObj.FullName)

				if testData.id == uuid.Nil {
					assert.NotEqual(test, uuid.Nil, userObj.AggregateRoot.Id, "UUID should be generated")
				} else {
					assert.Equal(test, testData.id, userObj.AggregateRoot.Id, "UUID should match the provided ID")
				}
			}
		})
	}
}
