package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/join")
	router.POST("/create", CreateGameRoute)

	return router
}
