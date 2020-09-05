package router

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/Jugendhackt/yujo-backend/config"
	"github.com/Jugendhackt/yujo-backend/models"
	"github.com/gin-gonic/gin"
)

var pinRandomMax = big.NewInt(999999 - 100000)

func CreateGameRoute(context *gin.Context) {
	randomPin, _ := rand.Int(rand.Reader, pinRandomMax)
	game := models.Game{
		GamePin: randomPin.Uint64(),
	}
	config.DB.Create(&game)

	context.JSON(http.StatusOK, game)
}

func JoinGameRoute(context *gin.Context) {

}
