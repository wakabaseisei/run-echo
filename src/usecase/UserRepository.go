package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type UserRepository interface {
	FindAll(db *gorm.DB) (users []domain.User, err error)
	FindByID(db *gorm.DB, id int) (post domain.User, err error)
	PostByForm(db *gorm.DB, sex int, introduction string) (post domain.User, err error)
	DeleteByID(db *gorm.DB, id int) (post domain.User, err error)
	UpdateByID(db *gorm.DB, id int, sex int, introduction string) (post domain.User, err error)
}
