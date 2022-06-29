package service

import (
	"testing"

	"github.com/sd01dev/nexign/internal/model"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFixSpelling_Success(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	checkerService := NewChecker(logger)

	entryText := model.TextToCheck{
		Texts: []string{"Ашибка", "Пагода"},
	}
	expectedResult := model.TextToCheck{
		Texts: []string{"Ошибка", "Погода"},
	}

	resultText, err := checkerService.FixSpelling(entryText)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, resultText)
}
