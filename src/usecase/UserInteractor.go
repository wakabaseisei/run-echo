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

func (interactor *UserInteractor) Post(sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	createdUser, err := interactor.UserRepository.PostByForm(db, sex, introduction)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = createdUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Delete(id int) (user domain.UserGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundUser, err := interactor.UserRepository.DeleteByID(db, id)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Update(id int, sex int, introduction string) (user domain.UserGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	updatedUser, err := interactor.UserRepository.UpdateByID(db, id, sex, introduction)
	if err != nil {
		return domain.UserGet{}, NewResultStatus(404, err)
	}

	user = updatedUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
