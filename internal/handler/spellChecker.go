package handler

import (
	"net/http"

	"github.com/sd01dev/nexign/internal/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const FixSpellingPath = "/fixSpelling"

type Checker struct {
	logger  *zap.Logger
	checker CheckerService
}

type CheckerService interface {
	FixSpelling(model.TextToCheck) (model.TextToCheck, error)
}

func NewChecker(logger *zap.Logger, checker CheckerService) *Checker {
	return &Checker{
		logger:  logger,
		checker: checker,
	}
}

func (ch *Checker) FixSpelling(c *gin.Context) {
	requestedText := model.TextToCheck{}
	if err := c.BindJSON(&requestedText); err != nil {
		ch.logger.Error("bindJSON fail", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	requestedText, err := ch.checker.FixSpelling(requestedText)
	if err != nil {
		ch.logger.Error("fixSpelling method fail", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, &requestedText)
}
