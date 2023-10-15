package logger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	logger := InitLogger()
	logger.Info().Msg("log info message")
}
