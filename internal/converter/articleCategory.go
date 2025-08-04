package converter

import (
	"zll.blog.com/internal/model"
	"zll.blog.com/internal/pb"
)

func ProtoToModelForArticleCategory(p *pb.ArticleCategory) *model.ArticleCategory {
	return &model.ArticleCategory{
		ArticleId: uint(p.ArticleId),
		CategoryId: uint(p.CategoryId),
	}
}

func ModelToProtoForArticleCategory(m *model.ArticleCategory) *pb.ArticleCategory {
	return &pb.ArticleCategory{
		ArticleId: uint64(m.ArticleId),
		CategoryId: uint64(m.CategoryId),
	}
}

func ProtoToModelForArticleCategoryList(p *pb.ArticleCategoryListResponse) []*model.ArticleCategory {
	var articleCategoryList []*model.ArticleCategory
	for _, ap := range p.ArticleCategoryList {
		articleCategoryList = append(articleCategoryList, ProtoToModelForArticleCategory(ap))
	}
	return articleCategoryList
}

// ModelToProto 将 Model 转换为 Proto
func ModelToProtoForArticleCategoryList(m []*model.ArticleCategory) *pb.ArticleCategoryListResponse {
	var articleCategoryList pb.ArticleCategoryListResponse
	for _, am := range m {
		articleCategoryList.ArticleCategoryList = append(articleCategoryList.ArticleCategoryList, ModelToProtoForArticleCategory(am))
	}
	return &articleCategoryList
}