package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
	"github.com/wakabaseisei/runapp/usecase"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) usecase.UserRepo {
	return &UserRepository{Conn: conn}
}

func (repo *UserRepository) FindByID(id int) (user domain.User, err error) {
	user = domain.User{}
	repo.Conn.First(&user, id)
	if user.Id <= 0 {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) FindAll() (users []domain.User, err error) {
	users = []domain.User{}
	repo.Conn.Find(&users)
	// if users <= 0 {
	// 	return domain.User{}, errors.New("user is not found")
	// }
	return users, nil
}

func (repo *UserRepository) PostByForm(sex int, introduction string) (user domain.User, err error) {
	user = domain.User{Sex: sex, Introduction: introduction}
	result := repo.Conn.Create(&user)
	if result.Error != nil {
		return domain.User{}, errors.New("user was not created")
	}

	return user, nil
}

func (repo *UserRepository) DeleteByID(id int) (user domain.User, err error) {
	user = domain.User{}
	result := repo.Conn.Delete(&user, id)
	if result.Error != nil {
		return domain.User{}, errors.New("user was not deleted")
	}
	return user, nil
}

func (repo *UserRepository) UpdateByID(id int, sex int, introduction string) (user domain.User, err error) {
	user = domain.User{}
	result := repo.Conn.Transaction(func(tx *gorm.DB) error {
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
