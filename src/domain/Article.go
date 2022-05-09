package domain

import "time"

type Article struct {
	Id        int
	title     Title
	content   Content
	IsPaid    bool
	userId    UserId
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(title string, coontent string, isPaid bool, userId int) (*Article, error) {
	t, err := newTitle(title)
	if err != nil {
		return nil, err
	}

	c, err := newContent(coontent)
	if err != nil {
		return nil, err
	}

	u, err := newUserId(userId)
	if err != nil {
		return nil, err
	}

	return &Article{title: *t, content: *c, userId: *u}, nil
}
