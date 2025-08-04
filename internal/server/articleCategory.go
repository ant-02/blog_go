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

type articleCategoryServer struct {
	pb.UnimplementedArticleCategoryServiceServer
	articleCategoryService service.ArticleCategoryService
}

func NewArticleCategoryServer(db *gorm.DB) *articleCategoryServer {
	articleCategoryRepo := repository.NewArticleCategoryRepository(db)
	articleCategoryService := service.NewArticleCategoryService(articleCategoryRepo)
	return &articleCategoryServer{articleCategoryService: articleCategoryService}
}

func (c *articleCategoryServer) GetByCategoryId(ctx context.Context, in *pb.ArticleCategoryIdRequest) (*pb.ArticleCategoryListResponse, error) {
	articleCategoryList, err := c.articleCategoryService.GetByCategoryId(uint(in.Id), int(in.Count))
	if err != nil {
		logger.Error("Failed to get articleCategoryList by categoryId", map[string]interface{}{
			"event": "get get articleCategoryList by categoryId",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success to get articleCategoryList by categoryId", map[string]interface{}{
		"event": "get get articleCategoryList by categoryId",
	})
	return converter.ModelToProtoForArticleCategoryList(articleCategoryList), err
}

func (c *articleCategoryServer) GetCategoryIdByArticleId(ctx context.Context, in *pb.ArticleCategoryIdRequest) (*pb.ArticleCategoryIdResponse, error) {
	id, err := c.articleCategoryService.GetCategoryIdByArticleId(uint(in.Id))
	return &pb.ArticleCategoryIdResponse{Id: uint64(id)}, err
}

func (c *articleCategoryServer) SaveArticleCategory(ctx context.Context, in *pb.ArticleCategoryRequest) (*pb.ArticleCategoryResponse, error) {
	ok, err := c.articleCategoryService.SaveArticleCategory(converter.ProtoToModelForArticleCategory(in.ArticleCategory))
	return &pb.ArticleCategoryResponse{Ok: ok}, err
}
