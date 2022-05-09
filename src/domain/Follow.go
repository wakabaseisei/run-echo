package domain

import "time"

type Follow struct {
	Id        int
	userId    UserId
	viewerId  ViewerId
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newFollow(userId int, viewerId int) (*Follow, error) {
	u, err := newUserId(userId)
	if err != nil {
		return nil, err
	}

	v, err := newViewerId(viewerId)
	if err != nil {
		return nil, err
	}

	return &Follow{userId: *u, viewerId: *v}, nil
}
