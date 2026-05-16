package handler

import (
	"fmt"
	"net/http"

	"github.com/Erick-VP/godin/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		RespondWithError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		RespondWithError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to delete opening with id: %s", id))
		return
	}

	RespondWithSuccess(ctx, "delete-opening", opening)
}
