package hw02unpackstring

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
	"unicode"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

var ErrInvalidString = errors.New("invalid string")

func init() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck
}

func Unpack(str string) (string, error) {
	runes := []rune(str)
	if len(runes) == 0 {
		return "", nil
	}

	var outStr strings.Builder

	leftRune := runes[0]
	if unicode.IsDigit(leftRune) {
		return "", ErrInvalidString
	}
	runesLen := len(runes)
	for i := 0; i < runesLen; i++ {
		r := runes[i]
		slog.Debug(i, " ", string(r))
		if i < runesLen-1 {
			rightRune := runes[i+1]
			if unicode.IsDigit(r) {
				slog.Debug(r, rightRune)
				return "", ErrInvalidString
			}
			if unicode.IsDigit(rightRune) {
				c, err := strconv.Atoi(string(rightRune))
				if err != nil {
					slog.Panic("что-то неправильно написано....")
				}
				outStr.WriteString(strings.Repeat(string(r), c))
				i++
			} else {
				outStr.WriteString(string(r))
			}
		} else {
			if !unicode.IsDigit(r) {
				outStr.WriteString(string(r))
			}
		}
	}

	return outStr.String(), nil
}
