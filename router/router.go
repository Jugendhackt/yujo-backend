package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/create", CreateGameRoute)
	router.GET("/join")

	return router
}
