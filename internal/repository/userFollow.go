package repository

import (
	"gorm.io/gorm"
	db "zll.blog.com/internal/database"
	"zll.blog.com/internal/model"
)

type userFollowRepository struct {
	db *gorm.DB
}

type UserFollowRepository interface {
	GetFollowerIds(id uint) ([]uint, error)
	GetFollowingIds(id uint) ([]uint, error)
	GetIsFollow(follower_id, following_id uint) (bool, error)
	SetIsFollow(follower_id, following_id uint, isFollow bool) error
}

func NewUserFollowRepository(db *gorm.DB) UserFollowRepository {
	return &userFollowRepository{db: db}
}

func (ufr *userFollowRepository) GetFollowerIds(id uint) ([]uint, error) {
	var followerIds []uint
	err := ufr.db.Model(&model.UserFollow{}).
		Select("follower_id").
		Where("is_deleted = '0' and following_id = ?", id).
		Find(&followerIds).Error
	return followerIds, err
}

func (ufr *userFollowRepository) GetFollowingIds(id uint) ([]uint, error) {
	var followingIds []uint
	err := ufr.db.Model(&model.UserFollow{}).
		Select("following_id").
		Where("is_deleted = '0' and follower_id = ?", id).
		Find(&followingIds).Error
	return followingIds, err
}

func (ufr *userFollowRepository) GetIsFollow(follower_id, following_id uint) (bool, error) {
	return db.Exists(ufr.db, &model.UserFollow{}, "follower_id = ? and following_id = ? and is_deleted = '0'", follower_id, following_id)
}

func (ufr *userFollowRepository) SetIsFollow(follower_id, following_id uint, isFollow bool) error {
	is, err := db.Exists(ufr.db, &model.UserFollow{}, "follower_id = ? and following_id = ?", follower_id, following_id)
	if err != nil {
		return err
	}
	if is {
		if err := ufr.db.Model(&model.UserFollow{}).
			Where("follower_id = ? and following_id = ?", follower_id, following_id).
			Update("is_deleted", !isFollow).Error; err != nil {
			return err
		}
	} else {
		if err := ufr.db.Create(&model.UserFollow{FollowerId: follower_id, FollowingId: following_id, IsDeleted: !isFollow}).Error; err != nil {
			return err
		}
	}
	return nil
}
