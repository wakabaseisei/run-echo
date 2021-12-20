package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type PostRepository struct{}

func (repo *PostRepository) FindByID(db *gorm.DB, id int) (post domain.Post, err error) {
	post = domain.Post{}
	db.First(&post, id)
	if post.Id <= 0 {
		return domain.Post{}, errors.New("post is not found")
	}
	return post, nil
}
