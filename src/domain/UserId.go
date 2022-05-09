package domain

import "errors"

type UserId struct {
	value int
}

func newUserId(v int) (*UserId, error) {
	if v < 0 {
		return nil, errors.New("THe userId cannot be smaller than 0")
	}
	return &UserId{value: v}, nil
}
