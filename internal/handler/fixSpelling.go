package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexign/internal/model"
	"nexign/internal/service"
)

func FixSpelling(c *gin.Context) {
	requestedText := model.TextToCheck{}
	if err := c.BindJSON(&requestedText); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	requestedText, err := service.FixSpelling(requestedText)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, &requestedText)
}
