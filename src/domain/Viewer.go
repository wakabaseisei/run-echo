package domain

import (
	"errors"
	"time"
)

type Viewer struct {
	Id           int
	sex          Sex
	introduction Introduction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func newViewer(sex int, intro string) (*Viewer, error) {
	s, err := newSex(sex)
	if err != nil {
		return nil, errors.New("The Sex field cannot be empty")
	}

	i, err := newIntroduction(intro)
	if err != nil {
		return nil, errors.New("The Introduction field cannot be empty")
	}

	return &Viewer{sex: *s, introduction: *i}, nil
}
