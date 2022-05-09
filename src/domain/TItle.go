package domain

import (
	"errors"
	"unicode/utf8"
)

type Title struct {
	value string
}

func newTitle(v string) (*Title, error) {
	if utf8.RuneCountInString(v) < 3 {
		return nil, errors.New("The title cannot be shorter than 3 charactors")
	}
	return &Title{value: v}, nil
}
