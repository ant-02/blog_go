package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/repository"
	"zll.blog.com/internal/service"
	"zll.blog.com/pkg/logger"
)

type categoryServer struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
}

func NewCategoryServer(db *gorm.DB) *categoryServer {
	categoryRepo := repository.NewCategoryRepository(*db)
	categoryService := service.NewCategoryService(categoryRepo)
	return &categoryServer{categoryService: categoryService}
}

func (s *categoryServer) GetAll(ctx context.Context, in *emptypb.Empty) (*pb.CategoryListResponse, error) {
	categories, err := s.categoryService.GetAll()
	if err != nil {
		logger.Error("Failed to get all categories", map[string]interface{}{
			"event": "get all categories",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success to get all categories", map[string]interface{}{
		"event": "get all categories",
	})
	return converter.ModelToProtoForCategoryList(categories), nil
}
