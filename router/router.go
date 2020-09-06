package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	router.POST("/create", CreateGameRoute)
	router.POST("/join/:gamePin", JoinGameRoute)
	router.GET("/game/:uuid", GetGameInfoRoute)

	router.GET("/game/:uuid/round/:id", GetQuestionRoute)
	router.POST("/game/:uuid/round/:id", SendAnswerRoute)

	router.GET("/", func(c *gin.Context) {
		c.FileFromFS("/app.html", pkger.Dir("/frontend"))
	})
	router.GET("/stylesheet.css", func(c *gin.Context) {
		c.FileFromFS("/stylesheet.css", pkger.Dir("/frontend"))
	})
	router.StaticFS("/script", pkger.Dir("/frontend/script"))
	router.StaticFS("/images", pkger.Dir("/frontend/images"))

	return router
}
