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

func (repo *UserRepository) FindAll(db *gorm.DB) (users []domain.User, err error) {
	users = []domain.User{}
	db.Find(&users)
	// if users <= 0 {
	// 	return domain.User{}, errors.New("user is not found")
	// }
	return users, nil
}

func (repo *UserRepository) PostByForm(db *gorm.DB, sex int, introduction string) (user domain.User, err error) {
	user = domain.User{Sex: sex, Introduction: introduction}
	result := db.Create(&user)
	if result.Error != nil {
		return domain.User{}, errors.New("user was not created")
	}

	return user, nil
}

func (repo *UserRepository) DeleteByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{}
	result := db.Delete(&user, id)
	if result.Error != nil {
		return domain.User{}, errors.New("user was not deleted")
	}
	return user, nil
}

func (repo *UserRepository) UpdateByID(db *gorm.DB, id int, sex int, introduction string) (user domain.User, err error) {
	user = domain.User{}
	result := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
			return err
		}
		user.Sex = sex
		user.Introduction = introduction
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return domain.User{}, errors.New("user was not deleted")
	}
	return user, nil
}
