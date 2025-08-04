package model

type UserFollow struct {
	FollowerId  uint `json:"followerId"`
	FollowingId uint `json:"followingId"`
	IsDeleted   bool `json:"isDeleted"`
}
