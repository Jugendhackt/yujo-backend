package router

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/Jugendhackt/yujo-backend/config"
	"github.com/Jugendhackt/yujo-backend/models"
	"github.com/gin-gonic/gin"
)

var pinRandomMax = big.NewInt(999999 - 100000)

func CreateGameRoute(context *gin.Context) {
	valid := false
	var randomPin *big.Int
	for !valid {
		randomPin, _ = rand.Int(rand.Reader, pinRandomMax)
		existing := config.DB.Where(&models.Game{GamePin: randomPin.Uint64()})
		if existing.RowsAffected == 0 {
			valid = true
		}
	}

	var payload models.CreateUser
	if err := context.BindJSON(&payload); err != nil {
		log.Println("Binding failed:", err)
	}

	game := models.Game{
		GamePin:     randomPin.Uint64(),
		CreatorName: payload.Name,
	}
	config.DB.Create(&game)
	if config.DB.Error != nil {
		context.AbortWithError(http.StatusInternalServerError, config.DB.Error)
		return
	}

	context.JSON(http.StatusOK, game)
}

func JoinGameRoute(context *gin.Context) {
	gamePin, err := strconv.ParseInt(context.Param("gamePin"), 10, 64)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var games []models.Game
	config.DB.Where(&models.Game{GamePin: uint64(gamePin)}).Find(&games)
	if len(games) == 0 {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var payload models.CreateUser
	if err := context.BindJSON(&payload); err != nil {
		log.Println("Binding failed:", err)
	}

	game := games[0]

	game.TeamMateName = payload.Name

	config.DB.Save(&game)

	if config.DB.Error != nil {
		context.AbortWithError(http.StatusInternalServerError, config.DB.Error)
		return
	}

	context.JSON(http.StatusOK, game)
}
