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
	if round.ID == 0 {
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

func SendAnswerRoute(context *gin.Context) {
	uuid := context.Param("uuid")
	roundID, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var round models.Round
	config.DB.Where(models.Round{GameID: uuid, GameBaseID: roundID}).First(&round)

	var payload jsonmodels.Answer
	if err := context.BindJSON(&payload); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	answer := models.Answer{
		User: payload.Answer,
	}
	round.Answers = append(round.Answers, answer)
	config.DB.Save(&round)
}

func GetFightResult(context *gin.Context) {
	uuid := context.Param("uuid")
	roundID, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var round models.Round
	config.DB.Joins("Answers").Where(models.Round{GameID: uuid, GameBaseID: roundID}).First(&round)

	if len(round.Answers) < 2 {
		context.AbortWithStatus(http.StatusConflict)
		return
	}

	var game models.Game
	config.DB.Joins("Enemy").Joins("Creator").Joins("TeamMate").Where(models.Game{ID: round.GameID}).First(&game)

	if round.Answers[0] == round.Answers[1] {
		game.Enemy.Healthpoints = game.Enemy.Healthpoints - 10
		game.Creator.Healthpoints = game.Creator.Healthpoints - 5
		game.TeamMate.Healthpoints = game.TeamMate.Healthpoints - 5
	} else {
		game.Creator.Healthpoints = game.Creator.Healthpoints - 10
		game.TeamMate.Healthpoints = game.TeamMate.Healthpoints - 10
	}

	config.DB.Save(&game)

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
		CorrectAnswer: round.Answers[0] == round.Answers[1],
		NextRoundID:   int(round.GameBaseID + 1),
	}

	context.JSON(http.StatusOK, response)
}
