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

type articleServer struct {
	pb.UnimplementedArticleServiceServer
	articleService service.ArticleService
}

func NewArticleServer(db *gorm.DB) *articleServer {
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	return &articleServer{articleService: articleService}
}

func (s *articleServer) GetArticleById(c context.Context, r *pb.ArticleIdRequest) (*pb.ArticleResponse, error) {
	article, err := s.articleService.GetArticleById(uint(r.Id))
	if err != nil {
		logger.Error("Failed", map[string]interface{}{
			"event": "get a article by id",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success", map[string]interface{}{
		"event": "get a article by id",
	})
	return &pb.ArticleResponse{Article: converter.ModelToProtoForArticle(&article)}, nil
}

func (s *articleServer) GetArticleDTOById(c context.Context, in *pb.ArticleIdRequest) (*pb.ArticleDTOResponse, error) {
	articleDTO, err := s.articleService.GetArticleDTOById(uint(in.Id))
	if err != nil {
		logger.Error("Failed", map[string]interface{}{
			"event": "get a articleDTO by id",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success", map[string]interface{}{
		"event": "get a articleDTO by id",
	})
	return converter.ModelToProtoForArticleDTO(articleDTO), nil
}

func (s *articleServer) GetArticleDTOsByKeyWords(ctx context.Context, in *pb.ArticleKeywordsRequest) (*pb.ArticleDTOsResponse, error) {
	articleDTOs, err := s.articleService.GetArticleDTOsByKeywords(in.Keywords)
	if err != nil {
		logger.Error("Failed", map[string]interface{}{
			"event": "get a articleDTOs by keywords",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success", map[string]interface{}{
		"event": "get a articleDTOs by keywords",
	})
	var aDs *pb.ArticleDTOsResponse = &pb.ArticleDTOsResponse{ArticleDTOs: []*pb.ArticleDTOResponse{}}
	for _, a := range articleDTOs {
		aDs.ArticleDTOs = append(aDs.ArticleDTOs, converter.ModelToProtoForArticleDTO(a))
	}
	return aDs, nil
}

func (s *articleServer) GetArticleDTOsByUserId(ctx context.Context, in *pb.ArticleIdRequest) (*pb.ArticleDTOsResponse, error) {
	articleDTOs, err := s.articleService.GetArticleDTOsByUserId(uint(in.Id), int(in.Page), int(in.PageSize), in.Status)
	if err != nil {
		logger.Error("Failed", map[string]interface{}{
			"event": "get a articleDTOs by userId",
			"error": err.Error(),
		})
		return nil, err
	}
	logger.Info("Success", map[string]interface{}{
		"event": "get a articleDTOs by userId",
	})
	var aDs *pb.ArticleDTOsResponse = &pb.ArticleDTOsResponse{ArticleDTOs: []*pb.ArticleDTOResponse{}, Count: uint64(articleDTOs.Count)}
	for _, a := range articleDTOs.ArticleDTOs {
		aDs.ArticleDTOs = append(aDs.ArticleDTOs, converter.ModelToProtoForArticleDTO(a))
	}
	return aDs, nil
}

func (s *articleServer) SaveArticle(ctx context.Context, in *pb.ArticleRequest) (*pb.SaveArticleResponse, error) {
	res, err := s.articleService.SaveArticle(converter.ProtoToModelForArticle(in.Article))
	return &pb.SaveArticleResponse{Res: res}, err
}
