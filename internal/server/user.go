package server

import (
	"context"

	"gorm.io/gorm"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/repository"
	"zll.blog.com/internal/service"
	"zll.blog.com/pkg/logger"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserServer(db *gorm.DB) *userServer {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	return &userServer{userService: userService}
}

func (us *userServer) GetUserDTOById(ctx context.Context, in *pb.UserIdRequest) (*pb.UserDTOResponse, error) {
	userDTO, err := us.userService.GetUserDTOById(uint(in.Id))
	return converter.ModelToProtoForUserDTO(userDTO), err
}

func (us *userServer) Login(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserResponse, error) {
	user, err := us.userService.Login(in.Phone, in.Password)
	return converter.ModelToProtoForUser(user), err
}

func (us *userServer) GetUserById(ctx context.Context, in *pb.UserIdRequest) (*pb.UserResponse, error) {
	user, err := us.userService.GetUserById(uint(in.Id))
	return converter.ModelToProtoForUser(user), err
}

func (us *userServer) GetUserDTOsByKeywords(ctx context.Context, in *pb.UserKeywordsRequest) (*pb.UserDTOsResponse, error) {
	userDTOs, err := us.userService.GetUserDTOsByKeywords(in.Keywords)
	if err != nil {
		logger.Error("Failed", map[string]interface{}{
			"event": "get a userDTOs by keywords",
			"error": err.Error(),
		})
		return nil, err
	}
	var uDs *pb.UserDTOsResponse = &pb.UserDTOsResponse{UserDTOs: []*pb.UserDTOResponse{}}
	for _, u := range userDTOs {
		uDs.UserDTOs = append(uDs.UserDTOs, converter.ModelToProtoForUserDTO(u))
	}
	return uDs, nil
}

func (us *userServer) Register(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserResponse, error) {
	user, err := us.userService.Register(in.Phone, in.Password)
	if err != nil {
		return nil, err
	}
	return converter.ModelToProtoForUser(user), err
}
