package controllers

import (
	"strconv"

	"github.com/wakabaseisei/runapp/interfaces/database"
	"github.com/wakabaseisei/runapp/usecase"
)

type PostsController struct {
	Interactor usecase.PostInteractor
}

func NewPostsController(db database.DB) *PostsController {
	return &PostsController{
		Interactor: usecase.PostInteractor{
			DB:             &database.DBRepository{DB: db},
			PostRepository: &database.PostRepository{},
		},
	}
}

func (controller *PostsController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	post, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", post))
}
