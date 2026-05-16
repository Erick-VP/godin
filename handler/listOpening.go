package handler

import (
	"net/http"

	"github.com/Erick-VP/godin/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, "error fetching openings")
		return
	}

	RespondWithSuccess(ctx, "list openings", openings)
}
