package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type ArticleRepository interface {
	FindByID(db *gorm.DB, id int) (article domain.Article, err error)
}
