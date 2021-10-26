package hw02unpackstring

import (
	"errors"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

var ErrInvalidString = errors.New("invalid string")

func init() {
	slog = logger.NewSugaredLogger()
}
