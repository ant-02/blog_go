package service

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/repository"
)

type UserService interface {
	GetUserDTOById(id uint) (*model.UserDTO, error)
	Login(phone string, password string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetUserDTOsByKeywords(keywords string) ([]*model.UserDTO, error)
} 

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepo: userRepository}
}

func (us *userService) GetUserDTOById(id uint) (*model.UserDTO, error) {
	return us.userRepo.GetUserDTOById(id)
}

func (us *userService) Login(phone string, password string) (*model.User, error) {
	return us.userRepo.Login(phone, password)
}

func (us *userService) GetUserById(id uint) (*model.User, error) {
	return us.userRepo.GetUserById(id)
} 

func (us *userService) GetUserDTOsByKeywords(keywords string) ([]*model.UserDTO, error) {
	return us.userRepo.GetUserDTOsByKeywords(keywords)
}