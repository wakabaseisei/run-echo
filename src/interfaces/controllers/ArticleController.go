package controllers

import (
	"strconv"

	"github.com/wakabaseisei/runapp/interfaces/database"
	"github.com/wakabaseisei/runapp/usecase"
)

type ArticlesController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticlesController(db database.DB) *ArticlesController {
	return &ArticlesController{
		Interactor: usecase.ArticleInteractor{
			DB:                &database.DBRepository{DB: db},
			ArticleRepository: &database.ArticleRepository{},
		},
	}
}

func (controller *ArticlesController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}
