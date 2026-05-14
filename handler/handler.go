package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET openings",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET opening",
	})
}
