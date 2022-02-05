package infrastructure

import (
	"time"

	"github.com/gin-contrib/cors"
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
	r.Gin.Use(cors.New(cors.Config{
		// AllowOrigins: []string{
		// 	"http://0.0.0.0:30090",
		// },
		AllowAllOrigins: true,
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Accept",
			"Referer",
			"User-Agent",
			"Origin",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
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
