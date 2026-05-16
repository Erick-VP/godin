package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RespondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"error_code": statusCode,
	})
}

func RespondWithSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("operation from handler: \"%s\" succeeded", op),
		"data":      data,
	})
}
