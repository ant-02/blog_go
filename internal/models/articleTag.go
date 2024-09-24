package models

type ArticleTag struct {
	ID        uint `json:"id"`
	TagId     uint `json:"tagId"`
	ArticleId uint `json:"articleId"`
}

type ArticlePreviewTagsRes struct {
	ArticlePreviewTags []ArticlePreviewTag `json:"articlePreviewTags"`
	Size               int                 `json:"size"`
}

type ArticlePreviewTag struct {
	ID              uint             `json:"id"`
	Name            string           `json:"name"`
	ArticlePreviews []ArticlePreview `json:"articlePreviews"`
}
