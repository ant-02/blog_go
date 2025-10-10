package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"gorm.io/gorm"
	"zll.blog.com/internal/model"
)

type UserRepository interface {
	GetUserDTOById(id uint) (*model.UserDTO, error)
	Login(phone string, password string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetUserDTOsByKeywords(keywords string) ([]*model.UserDTO, error)
	Register(phone string, password string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUserDTOById(id uint) (*model.UserDTO, error) {
	var userDTO *model.UserDTO
	err := ur.db.Model(&model.User{}).
		Select("id", "username", "avatar").
		Find(&userDTO, id).Error
	return userDTO, err
}

func (ur *userRepository) Login(phone string, password string) (*model.User, error) {
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])
	var user *model.User
	err := ur.db.Where(&model.User{Phone: phone, Password: password}).
		Find(&user).Error
	return user, err
}

func (ur *userRepository) GetUserById(id uint) (*model.User, error) {
	var user *model.User
	err := ur.db.Find(&user, id).Error
	return user, err
}

func (ur *userRepository) GetUserDTOsByKeywords(keywords string) ([]*model.UserDTO, error) {
	var userDTOs []*model.UserDTO
	keywords = "%" + keywords + "%"
	err := ur.db.Model(&model.User{}).
		Select("id", "username", "avatar").
		Where("username LIKE ?", keywords).
		Find(&userDTOs).Error
	return userDTOs, err
}

func (ur *userRepository) Register(phone string, password string) (*model.User, error) {
	hash := md5.Sum([]byte(password))
	ecd := hex.EncodeToString(hash[:])
	var user *model.User
	err := ur.db.Where(&model.User{Phone: phone}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Id != 0 {
		return nil, errors.New("phone have been created")
	}
	user.Phone = phone
	user.Password = ecd
	user.CreatedAt = time.Now()
	err = ur.db.Model(&model.User{}).
		Create(user).Error
	if err != nil {
		return nil, err
	}
	user.Password = password
	return user, nil
}
