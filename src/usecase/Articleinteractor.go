package usecase

import "github.com/wakabaseisei/runapp/domain"

type ArticleInteractor struct {
	DB                DBRepository
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) Get(id int) (article domain.ArticleGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()

	foundArticle, err := interactor.ArticleRepository.FindByID(db, id)
	if err != nil {
		return domain.ArticleGet{}, NewResultStatus(404, err)
	}

	article = foundArticle.BuildForGet()
	return article, NewResultStatus(200, nil)
}
