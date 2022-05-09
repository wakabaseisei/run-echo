package domain

import "errors"

type BillingPrice struct {
	value int
}

func newBillingPrice(v int) (*BillingPrice, error) {
	if v < 0 {
		return nil, errors.New("The billing price cannot be smaller than zero")
	}
	return &BillingPrice{value: v}, nil
}
