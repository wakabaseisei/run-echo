package domain

import "time"

// ユーザー
type User struct {
	Id           int
	Sex          int
	Introduction string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
