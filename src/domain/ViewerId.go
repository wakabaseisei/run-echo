package domain

import "errors"

type ViewerId struct {
	value int
}

func newViewerId(v int) (*ViewerId, error) {
	if v < 0 {
		return nil, errors.New("THe viewerId cannot be smaller than 0")
	}
	return &ViewerId{value: v}, nil
}
