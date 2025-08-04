package service

import "zll.blog.com/internal/repository"

type userFollowService struct {
	userFollowRepository repository.UserFollowRepository
}

type UserFollowService interface {
	GetFollowerIds(id uint) ([]uint, error)
	GetFollowingIds(id uint) ([]uint, error)
	GetIsFollow(follower_id, following_id uint) (bool, error)
	SetIsFollow(follower_id, following_id uint, isFollow bool) error
}

func NewUserFollowService(userFollowRepository repository.UserFollowRepository) UserFollowService {
	return &userFollowService{userFollowRepository: userFollowRepository}
}

func (ufs *userFollowService) GetFollowerIds(id uint) ([]uint, error) {
	return ufs.userFollowRepository.GetFollowerIds(id)
}

func (ufs *userFollowService) GetFollowingIds(id uint) ([]uint, error) {
	return ufs.userFollowRepository.GetFollowingIds(id)
}

func (ufs *userFollowService) GetIsFollow(follower_id, following_id uint) (bool, error) {
	return ufs.userFollowRepository.GetIsFollow(follower_id, following_id)
}

func (ufs *userFollowService) SetIsFollow(follower_id, following_id uint, isFollow bool) error {
	return ufs.userFollowRepository.SetIsFollow(follower_id, following_id, isFollow)
}
