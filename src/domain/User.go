package domain

import (
	"time"
)

// ユーザー
type User struct {
	Id              int
	sex             Sex
	introduction    Introduction
	IsForumNeedPaid bool
	billingPrice    BillingPrice
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// コンストラクタの定義
func NewUser(sex int, intro string, price int) (*User, error) {
	s, err := newSex(sex)
	if err != nil {
		return nil, err
	}

	i, err := newIntroduction(intro)
	if err != nil {
		return nil, err
	}

	b, err := newBillingPrice(price)
	if err != nil {
		return nil, err
	}

	return &User{sex: *s, introduction: *i, billingPrice: *b}, nil
}
