package server

import (
	"context"

	"gorm.io/gorm"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/repository"
	"zll.blog.com/internal/service"
)

type articleTagServer struct {
	pb.UnimplementedArticleTagServiceServer
	articleTagService service.ArticleTagService
}

func NewArticleTagServer(db *gorm.DB) *articleTagServer {
	articleTagRepository := repository.NewArticleTagRepository(db)
	articleTagService := service.NewArticleTagService(articleTagRepository)
	return &articleTagServer{articleTagService: articleTagService}
}

func (ts *articleTagServer) GetTagIdsByArticleId(ctx context.Context, in *pb.ArticleTagIdRequest) (*pb.ArticleTagIdsResponse, error) {

	ids, err := ts.articleTagService.GetIdsByArticleId(uint(in.Id))
	var res []uint64
	for _, id := range ids {
		res = append(res, uint64(id))
	}
	return &pb.ArticleTagIdsResponse{Ids: res}, err
}
