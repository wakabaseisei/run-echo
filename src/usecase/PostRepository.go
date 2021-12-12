package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type PostRepository interface {
	FindByID(db *gorm.DB, id int) (post domain.Post, err error)
}
