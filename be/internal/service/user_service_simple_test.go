package service

import (
	"errors"
	"testing"

	"milestone3/be/internal/dto"
	"milestone3/be/internal/entity"
	"milestone3/be/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser_Simple(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	req := dto.UserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
	mockRepo.EXPECT().GetById(gomock.Any()).Return(entity.Users{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "donor",
	}, nil)

	result, err := userService.CreateUser(req)

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "john@example.com", result.Email)
}

func TestUserService_GetUserById_Simple(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().GetById(1).Return(entity.Users{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "donor",
	}, nil)

	result, err := userService.GetUserById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "John Doe", result.Name)
}

func TestUserService_GetUserById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().GetById(999).Return(entity.Users{}, errors.New("user not found"))

	result, err := userService.GetUserById(999)

	assert.Error(t, err)
	assert.Empty(t, result)
}