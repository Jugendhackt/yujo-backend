package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/create", CreateGameRoute)
	router.POST("/join/:gamePin")

	return router
}
