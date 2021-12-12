package usecase

import "github.com/wakabaseisei/runapp/domain"

type PostInteractor struct {
	DB             DBRepository
	PostRepository PostRepository
}

func (interactor *PostInteractor) Get(id int) (post domain.PostGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundPost, err := interactor.PostRepository.FindByID(db, id)
	if err != nil {
		return domain.PostGet{}, NewResultStatus(404, err)
	}

	post = foundPost.BuildForGet()
	return post, NewResultStatus(200, nil)
}
