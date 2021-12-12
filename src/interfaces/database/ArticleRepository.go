package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type ArticleRepository struct{}

func (repo *ArticleRepository) FindByID(db *gorm.DB, id int) (article domain.Article, err error) {
	article = domain.Article{}
	db.First(&article, id)
	if article.ID <= 0 {
		return domain.Article{}, errors.New("article is not found")
	}
	return article, nil
}
