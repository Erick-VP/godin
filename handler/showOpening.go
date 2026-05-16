package handler

import (
	"net/http"

	"github.com/Erick-VP/godin/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		RespondWithError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		RespondWithError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	RespondWithSuccess(ctx, "show opening", opening)
}
