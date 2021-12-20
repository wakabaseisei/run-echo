package controllers

import (
	"strconv"

	"github.com/wakabaseisei/runapp/interfaces/database"
	"github.com/wakabaseisei/runapp/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(db database.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			DB:             &database.DBRepository{DB: db},
			UserRepository: &database.UserRepository{},
		},
	}
}

func (controller *UserController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}
