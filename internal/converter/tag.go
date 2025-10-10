package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

func ProtoToModelForTagList(p *pb.TagListResponse) []*model.Tag {
	var tags []*model.Tag
	for _, ap := range p.Tags {
		tags = append(tags, &model.Tag{
			BaseModel: model.BaseModel{
				Id:        ap.Id,
				CreatedAt: ap.CreatedAt.AsTime(),
				UpdatedAt: ap.UpdateAt.AsTime(),
				DeletedAt: gorm.DeletedAt{Time: ap.DeleteAt.AsTime()},
			},
			Name: ap.Name,
		})
	}
	return tags
}

// ModelToProto 将 Model 转换为 Proto
func ModelToProtoForTagList(m []*model.Tag) *pb.TagListResponse {
	var tags pb.TagListResponse
	for _, am := range m {
		tags.Tags = append(tags.Tags, &pb.TagResponse{
			Id:        am.Id,
			Name:      am.Name,
			CreatedAt: timestamppb.New(am.CreatedAt),
			UpdateAt:  timestamppb.New(am.UpdatedAt),
			DeleteAt:  timestamppb.New(am.DeletedAt.Time),
		})
	}
	return &tags
}
