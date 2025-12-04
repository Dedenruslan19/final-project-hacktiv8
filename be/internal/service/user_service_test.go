package service

import (
	"errors"
	"testing"

	"milestone3/be/internal/dto"
	"milestone3/be/internal/entity"
	"milestone3/be/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name    string
		req     dto.UserRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "successful user creation",
			req: dto.UserRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func() {
				mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
				mockRepo.EXPECT().GetById(gomock.Any()).Return(entity.Users{
					Id:    1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  "donor",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "repository create error",
			req: dto.UserRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func() {
				mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name: "get user by id error after creation",
			req: dto.UserRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func() {
				mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
				mockRepo.EXPECT().GetById(gomock.Any()).Return(entity.Users{}, errors.New("user not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := userService.CreateUser(tt.req)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "Test User", result.Name)
				assert.Equal(t, "test@example.com", result.Email)
			}
		})
	}
}

func TestUserService_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name    string
		id      int
		setup   func()
		wantErr bool
	}{
		{
			name: "successful get user by id",
			id:   1,
			setup: func() {
				user := entity.Users{
					Id:    1,
					Name:  "Test User",
					Email: "test@example.com",
					Role:  "donor",
				}
				mockRepo.EXPECT().GetById(1).Return(user, nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			id:   999,
			setup: func() {
				mockRepo.EXPECT().GetById(999).Return(entity.Users{}, errors.New("user not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			result, err := userService.GetUserById(tt.id)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, result.Id)
				assert.Equal(t, "Test User", result.Name)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	// Create a hashed password for testing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	tests := []struct {
		name     string
		email    string
		password string
		setup    func()
		wantErr  bool
	}{
		{
			name:     "user not found",
			email:    "notfound@example.com",
			password: "password123",
			setup: func() {
				mockRepo.EXPECT().GetByEmail("notfound@example.com").Return(entity.Users{}, errors.New("user not found"))
			},
			wantErr: true,
		},
		{
			name:     "wrong password",
			email:    "test@example.com",
			password: "wrongpassword",
			setup: func() {
				user := entity.Users{
					Id:       1,
					Name:     "Test User",
					Email:    "test@example.com",
					Password: string(hashedPassword),
					Role:     "donor",
				}
				mockRepo.EXPECT().GetByEmail("test@example.com").Return(user, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			
			token, err := userService.GetUserByEmail(tt.email, tt.password)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}