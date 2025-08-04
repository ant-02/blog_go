package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

func ProtoToModelForArticle(p *pb.Article) *model.Article {
	return &model.Article{
		BaseModel: model.BaseModel{
			Id:        p.Id,
			CreatedAt: p.CreatedAt.AsTime(),
			UpdatedAt: p.UpdateAt.AsTime(),
			DeletedAt: p.DeleteAt.AsTime(),
		},
		Title:      p.Title,
		Summary:    p.Summary,
		Content:    p.Content,
		CoverImage: p.CoverImage,
		AuthorId:   p.AuthorId,
		Status:     p.Status,
	}
}

func ModelToProtoForArticle(m *model.Article) *pb.Article {
	return &pb.Article{
		Id:         uint64(m.Id),
		Title:      m.Title,
		Summary:    m.Summary,
		Content:    m.Content,
		CoverImage: m.CoverImage,
		AuthorId:   uint64(m.AuthorId),
		Status:     m.Status,
		CreatedAt:  timestamppb.New(m.CreatedAt),
		UpdateAt:   timestamppb.New(m.UpdatedAt),
		DeleteAt:   timestamppb.New(m.DeletedAt),
	}
}

func ProtoToModelForArticleDTO(p *pb.ArticleDTOResponse) *model.ArticleDTO {
	return &model.ArticleDTO{
		Id:         p.Id,
		Title:      p.Title,
		UpdatedAt:  p.UpdatedAt.AsTime(),
		CoverImage: p.CoverImage,
		Summary:    p.Summary,
	}
}

func ModelToProtoForArticleDTO(m *model.ArticleDTO) *pb.ArticleDTOResponse {
	return &pb.ArticleDTOResponse{
		Id:         uint64(m.Id),
		Title:      m.Title,
		UpdatedAt:  timestamppb.New(m.UpdatedAt),
		CoverImage: m.CoverImage,
		Summary:    m.Summary,
	}
}
