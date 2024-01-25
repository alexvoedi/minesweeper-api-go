package health

import "github.com/gin-gonic/gin"

const (
	SUCCESS = "success"
	ERROR   = "error"
)

func InitHealthRouter(router *gin.Engine) {
	health := router.Group("/health")
	{
		health.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": SUCCESS,
			})
		})
	}
}
