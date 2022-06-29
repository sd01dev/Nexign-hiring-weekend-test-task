package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sd01dev/nexign/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFixSpelling_BadRequest(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	checkerService := service.NewChecker(logger)
	checkerHandler := NewChecker(logger, checkerService)

	router := gin.Default()
	router.POST(FixSpellingPath, checkerHandler.FixSpelling)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, FixSpellingPath, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}
