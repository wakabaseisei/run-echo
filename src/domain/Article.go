package domain

// 記事
type Article struct {
	ID           int
	Title        string
	Content      string
	ThumbnailUrl *string
	IsPaid       bool
	PublishDate  string
	CreatedAt    int64
	UpdatedAt    int64
}

// 返す記事情報
type ArticleGet struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	ThumbnailUrl *string `json:"thumbnailUrl"`
	IsPaid       bool    `json:"isPaid"`
	PublishDate  string  `json:"publishDate"`
}

func (a *Article) BuildForGet() ArticleGet {
	article := ArticleGet{}
	article.ID = a.ID
	article.Title = a.Title
	article.IsPaid = a.IsPaid
	article.PublishDate = a.PublishDate
	if a.ThumbnailUrl != nil {
		article.ThumbnailUrl = a.ThumbnailUrl
	} else {
		empty := ""
		article.ThumbnailUrl = &empty
	}
	return article
}
