package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

func ModelToProtoForUserDTO(m *model.UserDTO) *pb.UserDTOResponse {
	return &pb.UserDTOResponse{
		Id:       uint64(m.Id),
		Username: m.Username,
		Avatar:   m.Avatar,
	}
}

func ProtoToModelForUserDTO(p *pb.UserDTOResponse) *model.UserDTO {
	return &model.UserDTO{
		Id:       uint(p.Id),
		Username: p.Username,
		Avatar:   p.Avatar,
	}
}

func ModelToProtoForUser(m *model.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        uint64(m.Id),
		Username:  m.Username,
		Avatar:    m.Avatar,
		Bio:       m.Bio,
		Phone:     m.Phone,
		Password:  m.Password,
		Email:     m.Email,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdateAt:  timestamppb.New(m.UpdatedAt),
		DeleteAt:  timestamppb.New(m.DeletedAt.Time),
	}
}

func ProtoToModelForUser(p *pb.UserResponse) *model.User {
	return &model.User{
		BaseModel: model.BaseModel{
			Id:        p.Id,
			CreatedAt: p.CreatedAt.AsTime(),
			UpdatedAt: p.UpdateAt.AsTime(),
			DeletedAt: gorm.DeletedAt{Time: p.DeleteAt.AsTime()},
		},
		Username: p.Username,
		Avatar:   p.Avatar,
		Phone:    p.Phone,
		Password: p.Password,
		Bio:      p.Bio,
		Email:    p.Email,
	}
}
