package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"zll.blog.com/internal/converter"
	"zll.blog.com/internal/pb"
	"zll.blog.com/internal/repository"
	"zll.blog.com/internal/service"
)

type tagServer struct {
	pb.UnimplementedTagServiceServer
	tagService service.TagService
}

func NewTagServer(db *gorm.DB) *tagServer {
	tagRepository := repository.NewTagRepository(db)
	tagService := service.NewTagService(tagRepository)
	return &tagServer{tagService: tagService}
}

func (ts *tagServer) GetAll(cxt context.Context, in *emptypb.Empty) (*pb.TagListResponse, error) {
	tags, err := ts.tagService.GetAll()
	return converter.ModelToProtoForTagList(tags), err
}
