package router

import (
	"net/http"

	"github.com/Jugendhackt/yujo-backend/config"
	jsonmodels "github.com/Jugendhackt/yujo-backend/json-models"
	"github.com/Jugendhackt/yujo-backend/models"
	"github.com/gin-gonic/gin"
)

func GetGameInfoRoute(context *gin.Context) {
	uuid := context.Param("uuid")

	var game models.Game
	config.DB.Where(&models.Game{ID: uuid}).Joins("Creator").Joins("TeamMate").First(&game)

	if game.TeamMateJoined != true {
		context.AbortWithStatus(http.StatusConflict)
		return
	}

	response := jsonmodels.GameInfo{
		Names: jsonmodels.PlayerNames{
			CreatorName:  game.Creator.Name,
			TeamMateName: game.TeamMate.Name,
		},
		HealthPoints: jsonmodels.HealthPoints{
			Creator:  game.Creator.Healthpoints,
			TeamMate: game.TeamMate.Healthpoints,
			Enemy:    game.Enemy.Healthpoints,
		},
	}

	context.JSON(http.StatusOK, response)
}
