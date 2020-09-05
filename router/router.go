package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	router.POST("/create", CreateGameRoute)
	router.POST("/join/:gamePin", JoinGameRoute)
	router.GET("/game/:uuid", GetGameInfoRoute)

	return router
}
