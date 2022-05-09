package domain

import (
	"errors"
	"unicode/utf8"
)

type Content struct {
	value string
}

func newContent(v string) (*Content, error) {
	if utf8.RuneCountInString(v) < 5 {
		return nil, errors.New("The content cannot be shorter than 5 charactors")
	}
	return &Content{value: v}, nil
}
