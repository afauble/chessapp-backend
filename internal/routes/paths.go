package routes

import (
	"github.com/gin-gonic/gin"
)

func Gin_url_setup(port string) {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.GET("/getLegalMoves")

	r.POST("/game/start")
	r.POST("/game/move")

	r.Run(":" + port)
}
