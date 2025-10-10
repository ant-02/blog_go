package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

func ProtoToModelForCategoryList(p *pb.CategoryListResponse) []*model.Category {
	var categories []*model.Category
	for _, ap := range p.Categories {
		categories = append(categories, &model.Category{
			BaseModel: model.BaseModel{
				Id:        ap.Id,
				CreatedAt: ap.CreatedAt.AsTime(),
				UpdatedAt: ap.UpdateAt.AsTime(),
				DeletedAt: gorm.DeletedAt{Time: ap.DeleteAt.AsTime()},
			},
			Name: ap.Name,
		})
	}
	return categories
}

// ModelToProto 将 Model 转换为 Proto
func ModelToProtoForCategoryList(m []*model.Category) *pb.CategoryListResponse {
	var categories pb.CategoryListResponse
	for _, am := range m {
		categories.Categories = append(categories.Categories, &pb.CategoryResponse{
			Id:        am.Id,
			Name:      am.Name,
			CreatedAt: timestamppb.New(am.CreatedAt),
			UpdateAt:  timestamppb.New(am.UpdatedAt),
			DeleteAt:  timestamppb.New(am.DeletedAt.Time),
		})
	}
	return &categories
}
