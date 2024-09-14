package services

import (
	"github.com/CeyhunBoran/shaffra-casestudy/internal/models"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user models.User) (*models.User, error) {
	return us.userRepository.CreateUser(user)
}

func (us *UserService) GetUser(userID string) (*models.User, error) {
	return us.userRepository.GetUser(userID)
}

func (us *UserService) UpdateUser(userID string, user models.User) (*models.User, error) {
	return us.userRepository.UpdateUser(userID, user)
}

func (us *UserService) DeleteUser(userID string) error {
	return us.userRepository.DeleteUser(userID)
}
