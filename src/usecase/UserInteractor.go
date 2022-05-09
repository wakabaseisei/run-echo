package usecase

import (
	"github.com/wakabaseisei/runapp/domain"
)

type UserUsecase interface {
	Get(id int) (user domain.User, resultStatus *ResultStatus)
	GetAll() (usersGet []domain.User, resultStatus *ResultStatus)
	Post(sex int, introduction string) (user domain.User, resultStatus *ResultStatus)
	Delete(id int) (user domain.User, resultStatus *ResultStatus)
	Update(id int, sex int, introduction string) (user domain.User, resultStatus *ResultStatus)
}

type UserInteractor struct {
	UserRepository UserRepo
}

func NewUserUsecase(userRepo UserRepo) UserUsecase {
	return &UserInteractor{UserRepository: userRepo}
}

func (interactor *UserInteractor) Get(id int) (user domain.User, resultStatus *ResultStatus) {

	user, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return domain.User{}, NewResultStatus(404, err)
	}

	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) GetAll() (usersGet []domain.User, resultStatus *ResultStatus) {

	users, err := interactor.UserRepository.FindAll()
	if err != nil {
		return []domain.User{}, NewResultStatus(404, err)
	}

	return users, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Post(sex int, introduction string) (user domain.User, resultStatus *ResultStatus) {
	// TODO: ここでNewUserしてエラー検出後に、userを作成したものをレポジトリに渡す

	user, err := interactor.UserRepository.PostByForm(sex, introduction)
	if err != nil {
		return domain.User{}, NewResultStatus(404, err)
	}

	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Delete(id int) (user domain.User, resultStatus *ResultStatus) {

	user, err := interactor.UserRepository.DeleteByID(id)
	if err != nil {
		return domain.User{}, NewResultStatus(404, err)
	}

	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Update(id int, sex int, introduction string) (user domain.User, resultStatus *ResultStatus) {

	user, err := interactor.UserRepository.UpdateByID(id, sex, introduction)
	if err != nil {
		return domain.User{}, NewResultStatus(404, err)
	}

	return user, NewResultStatus(200, nil)
}
