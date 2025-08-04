package server

import (
	"context"

	"gorm.io/gorm"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/repository"
	"zll.blog.com/internal/service"
)

type userFollowServer struct {
	pb.UnimplementedUserFollowServiceServer
	userFollowService service.UserFollowService
}

func NewUserFollowServer(db *gorm.DB) *userFollowServer {
	userFollowRepository := repository.NewUserFollowRepository(db)
	userFollowService := service.NewUserFollowService(userFollowRepository)
	return &userFollowServer{userFollowService: userFollowService}
}

func (ufs *userFollowServer) GetUserFollowerIds(ctx context.Context, in *pb.UserFollowIdRequest) (*pb.UserFollowIdResponse, error) {
	ids, err := ufs.userFollowService.GetFollowerIds(uint(in.Id))
	followerIds := make([]uint64, len(ids))
	for i, v := range ids {
		followerIds[i] = uint64(v)
	}
	return &pb.UserFollowIdResponse{Id: followerIds}, err
}

func (ufs *userFollowServer) GetUserFollowingIds(ctx context.Context, in *pb.UserFollowIdRequest) (*pb.UserFollowIdResponse, error) {
	ids, err := ufs.userFollowService.GetFollowingIds(uint(in.Id))
	followingIds := make([]uint64, len(ids))
	for i, v := range ids {
		followingIds[i] = uint64(v)
	}
	return &pb.UserFollowIdResponse{Id: followingIds}, err
}

func (ufs *userFollowServer) GetIsFollow(ctx context.Context, in *pb.IsFollowedRequest) (*pb.IsFollowResponse, error) {
	isFollow, err := ufs.userFollowService.GetIsFollow(uint(in.FollowerId), uint(in.FollowingId))
	return &pb.IsFollowResponse{IsFollow: isFollow}, err
}

func (ufs *userFollowServer) SetIsFollow(ctx context.Context, in *pb.IsFollowedRequest) (*pb.IsFollowResponse, error) {
	err := ufs.userFollowService.SetIsFollow(uint(in.FollowerId), uint(in.FollowingId), in.IsFollow)
	if err != nil {
		return &pb.IsFollowResponse{IsFollow: false}, err
	}
	return &pb.IsFollowResponse{IsFollow: true}, nil
}
