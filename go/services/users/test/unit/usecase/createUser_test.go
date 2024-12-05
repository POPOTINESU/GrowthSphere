package usecase_test

import (
	mocks "GROWTHSPHERE/services/users/test/mocks/repository"
	"GROWTHSPHERE/services/users/usecase"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestICreateUserUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrl)
	usecase := usecase.NewCreateUserUsecase(mockRepo)

	tests := []struct {
		name         string
		username     string
		accountName  string
		mockSetup    func()
		expectedID   uuid.UUID
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:        "Happy Path: valid inputs",
			username:    "valid_user",
			accountName: "valid_account",
			mockSetup: func() {
				userID := uuid.MustParse("49bbe7c6-7fcb-4dc9-8099-260462a29c5f")
				mockRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(userID, nil)
			},
			expectedID:   uuid.MustParse("49bbe7c6-7fcb-4dc9-8099-260462a29c5f"),
			expectErr:    false,
			expectErrMsg: "",
		},
		{
			name:        "Error: invalid username",
			username:    "",
			accountName: "valid_account",
			mockSetup: func() {
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedID:   uuid.Nil,
			expectErr:    true,
			expectErrMsg: "username cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			resultID, err := usecase.Execute(ctx, tt.username, tt.accountName)

			assert.Equal(t, tt.expectedID, resultID)
			if tt.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectErrMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
