package api

import "github.com/gin-gonic/gin"

func Gin_url_setup() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.POST("/game/start")
	r.POST("/game/move")

	r.Run(":8080")
}
