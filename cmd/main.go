package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nexign/internal/handler"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	router := gin.Default()
	router.POST("/fixSpelling", handler.FixSpelling)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
