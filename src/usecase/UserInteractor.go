package usecase

import (
	"github.com/wakabaseisei/runapp/domain"
)

type UserUsecase interface {
	Get(id int) (user domain.UserGet, resultStatus *ResultStatus)
	GetAll() (usersGet []domain.UserGet, resultStatus *ResultStatus)
	Post(sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus)
	Delete(id int) (user domain.UserGet, resultStatus *ResultStatus)
	Update(id int, sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus)
}

type UserInteractor struct {
	UserRepository UserRepo
}

func NewUserUsecase(userRepo UserRepo) UserUsecase {
	return &UserInteractor{UserRepository: userRepo}
}

func (interactor *UserInteractor) Get(id int) (user domain.UserGet, resultStatus *ResultStatus) {

	foundUser, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) GetAll() (usersGet []domain.UserGet, resultStatus *ResultStatus) {

	foundUsers, err := interactor.UserRepository.FindAll()
	if err != nil {
		return []domain.UserGet{}, NewResultStatus(404, err)
	}
	usersGet = []domain.UserGet{}

	for i := range foundUsers {
		usersGet = append(usersGet, foundUsers[i].BuildForGet())
	}
	return usersGet, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Post(sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus) {

	createdUser, err := interactor.UserRepository.PostByForm(sex, introduction)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = createdUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Delete(id int) (user domain.UserGet, resultStatus *ResultStatus) {

	foundUser, err := interactor.UserRepository.DeleteByID(id)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Update(id int, sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus) {

	updatedUser, err := interactor.UserRepository.UpdateByID(id, sex, introduction)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = updatedUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
