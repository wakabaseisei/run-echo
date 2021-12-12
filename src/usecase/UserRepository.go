package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (post domain.User, err error)
}
