package usecase

import (
	"github.com/wakabaseisei/runapp/domain"
)

type UserRepo interface {
	FindByID(id int) (user domain.User, err error)
	FindAll() (users []domain.User, err error)
	PostByForm(sex int, introduction string) (user domain.User, err error)
	DeleteByID(id int) (user domain.User, err error)
	UpdateByID(id int, sex int, introduction string) (user domain.User, err error)
}
