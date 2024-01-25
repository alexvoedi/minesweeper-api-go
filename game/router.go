package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGameRouter(router *gin.Engine) {
	game := router.Group("/games")
	{
		game.PUT("/", func(c *gin.Context) {
			var body CreateGameDto

			if err := c.BindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{
				"status": "success",
				"data":   CreateGame(body),
			})
		})
	}
}
