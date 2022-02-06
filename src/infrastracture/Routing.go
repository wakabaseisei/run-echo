package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/wakabaseisei/runapp/interfaces/controllers"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	postController := controllers.NewPostsController(r.DB)
	userController := controllers.NewUserController(r.DB)
	r.Gin.GET("/users", func(c *gin.Context) { userController.GetAll(c) })
	r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
	r.Gin.POST("/users", func(c *gin.Context) { userController.Post(c) })
	r.Gin.DELETE("/users/:id", func(c *gin.Context) { userController.Delete(c) })
	r.Gin.PUT("/users/:id", func(c *gin.Context) { userController.Update(c) })
	r.Gin.GET("/posts/:id", func(c *gin.Context) { postController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
