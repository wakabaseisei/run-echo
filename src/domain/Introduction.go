package domain

import "errors"

type Introduction struct {
	value string
}

func newIntroduction(v string) (*Introduction, error) {
	if v == "" {
		return nil, errors.New("Introduction field is not empty. Please fill this feild")
	}
	return &Introduction{value: v}, nil
}
