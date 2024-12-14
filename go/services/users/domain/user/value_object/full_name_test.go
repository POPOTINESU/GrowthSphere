package valueobject

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFullName(test *testing.T) {
	data := []struct {
		name      string
		firstName string
		lastName  string
		expected  string
		errMsg    string
	}{
		{
			name:      "valid full name",
			firstName: "test",
			lastName:  "user",
			expected:  "test user",
			errMsg:    "",
		},
		{
			name:      "first name empty",
			firstName: "",
			lastName:  "user",
			expected:  "",
			errMsg:    fmt.Sprintf("first name must be greater than or equal to %d", NameMinLength),
		},
		{
			name:      "last name empty",
			firstName: "test",
			lastName:  "",
			expected:  "",
			errMsg:    fmt.Sprintf("last name must be greater than or equal to %d", NameMinLength),
		},
		{
			name:      "first name exceeds max length",
			firstName: strings.Repeat("a", NameMaxLength+1),
			lastName:  "user",
			expected:  "",
			errMsg:    fmt.Sprintf("first name must be less than or equal to %d", NameMaxLength),
		},
		{
			name:      "last name exceeds max length",
			firstName: "test",
			lastName:  strings.Repeat("a", NameMaxLength+1),
			expected:  "",
			errMsg:    fmt.Sprintf("last name must be less than or equal to %d", NameMaxLength),
		},
	}

	for _, testData := range data {
		test.Run(testData.name, func(test *testing.T) {
			fullName, err := NewFullName(testData.firstName, testData.lastName)

			if testData.errMsg != "" {
				assert.Error(test, err)
				assert.Equal(test, testData.errMsg, err.Error(), "error message should match")
			} else {
				assert.NoError(test, err)
				assert.Equal(test, testData.expected, fullName.String(), "fullname should match expected")
			}
		})
	}
}
