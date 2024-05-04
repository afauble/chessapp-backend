package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStartingPosition(ctx *gin.Context) {
	var results map[string]string = make(map[string]string)

	ctx.IndentedJSON(http.StatusOK, results)
}
