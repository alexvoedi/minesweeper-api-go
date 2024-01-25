package routers

import (
	"github.com/alexvoedi/minesweeper-api-go/game"
	"github.com/alexvoedi/minesweeper-api-go/health"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())

	health.InitHealthRouter(router)
	game.InitGameRouter(router)

	return router
}
