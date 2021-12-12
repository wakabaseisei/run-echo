package usecase

import "github.com/wakabaseisei/runapp/domain"

type UserInteractor struct {
	DB             DBRepository
	UserRepository UserRepository
}

func (interactor *UserInteractor) Get(id int) (user domain.UserGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundUser, err := interactor.UserRepository.FindByID(db, id)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
