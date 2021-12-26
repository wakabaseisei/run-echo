package domain

import "time"

// 記事
type Post struct {
	Id           int
	Title        string
	Content      string
	ThumbnailUrl *string
	PublishDate  int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 返す記事情報
type PostGet struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	ThumbnailUrl *string `json:"thumbnailUrl"`
	PublishDate  int     `json:"publishDate"`
}

func (a *Post) BuildForGet() PostGet {
	post := PostGet{}
	post.Id = a.Id
	post.Title = a.Title
	post.Content = a.Content
	post.PublishDate = a.PublishDate
	if a.ThumbnailUrl != nil {
		post.ThumbnailUrl = a.ThumbnailUrl
	} else {
		empty := ""
		post.ThumbnailUrl = &empty
	}
	return post
}
