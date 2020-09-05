package router

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

func GetQuestionRoute(context *gin.Context) {
	uuid := context.Param("uuid")
	roundID, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var round models.Round
	config.DB.Where("id = ? AND game_id = ?", roundID, uuid).First(&round)

	var question models.Question
	if round == (models.Round{}) {
		// TODO: Generate new Round here
		var count int64
		config.DB.Table("questions").Count(&count)
		rand.Seed(time.Now().UnixNano())
		config.DB.Offset(rand.Intn(int(count - 1))).Limit(1).Find(&question)
		round = models.Round{
			GameBaseID: roundID,
			GameID:     uuid,
			QuestionID: question.ID,
		}
		config.DB.Save(&round)
	} else {
		config.DB.Where("ID = ?", round.QuestionID).First(&question)
	}

	response := jsonmodels.Question{
		Text: question.Text,
	}

	context.JSON(http.StatusOK, response)
}
