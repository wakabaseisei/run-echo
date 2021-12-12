package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
)

type UserRepository struct{}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{}
	db.First(&user, id)
	if user.Id <= 0 {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}
