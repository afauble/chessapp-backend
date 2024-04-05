package routes

import (
	"afauble/go/chessapp/internal/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Gin_url_setup(port string) {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.GET("/getLegalMoves/:index", func(ctx *gin.Context) {
		var testMoveMap map[string]string = make(map[string]string)
		indexStr := ctx.Param("index")
		index, _ := strconv.ParseInt(indexStr, 10, 8)
		knightMoves := strconv.FormatInt(int64(helpers.GetStandardKnightMoves(int8(index))), 2)
		rookMoves := strconv.FormatInt(int64(helpers.GetStandardRookMoves(int8(index))), 2)
		bishopMoves := strconv.FormatInt(int64(helpers.GetStandardBishopMoves(int8(index))), 2)
		testMoveMap["knight"] = knightMoves
		testMoveMap["rook"] = rookMoves
		testMoveMap["bishop"] = bishopMoves

		ctx.IndentedJSON(http.StatusOK, testMoveMap)
	})

	r.POST("/game/start")
	r.POST("/game/move")

	r.Run(":" + port)
}
