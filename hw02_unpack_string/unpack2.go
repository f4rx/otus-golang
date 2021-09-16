package hw02unpackstring

import (

	// "fmt".
	"strconv"
	"strings"
	"unicode"
)

type customChar struct {
	char    rune
	isDigit bool
}

func parseString(str string) ([]customChar, error) {
	chars := make([]customChar, 0)
	backslash := false

	// По правилам слеш может экранировать только число или слеш,
	// таким образом сам по себе не может существовать
	runeList := []rune(str)
	if len(runeList) == 1 && runeList[0] == '\\' {
		return make([]customChar, 0), ErrInvalidString
	}

	for i := 0; i < len(runeList); i++ {
		c := runeList[i]

		if !backslash && c == '\\' {
			backslash = true
			continue
		}

		// Можно экранировать только слеш и число, \n экранировать нельзя
		isValidedEscapedChar := func(r rune) bool {
			switch {
			case unicode.IsDigit(r):
				return true
			case r == '\\':
				return true
			default:
				return false
			}
		}
		if backslash && !isValidedEscapedChar(c) {
			return make([]customChar, 0), ErrInvalidString
		}

		chars = append(chars, customChar{c, !backslash && unicode.IsDigit(c)})
		backslash = false
	}
	return chars, nil
}

func UnpackStareks(str string) (string, error) {
	customChars, err := parseString(str)
	if err != nil {
		return "", err
	}

	if len(customChars) == 0 {
		return "", nil
	}

	if customChars[0].isDigit {
		return "", ErrInvalidString
	}

	var outStr strings.Builder
	runesLen := len(customChars)

	for i := 0; i < runesLen; i++ {
		r := customChars[i]
		outMessage := string(r.char)
		slog.Debug(i, " ", outMessage)

		if r.isDigit {
			// slog.Debug(r, rightRune)
			return "", ErrInvalidString
		}

		if i < runesLen-1 {
			rightChar := customChars[i+1]

			if rightChar.isDigit {
				repeatCount, err := strconv.Atoi(string(rightChar.char))
				if err != nil {
					slog.Panic("что-то неправильно написано....")
				}
				outMessage = strings.Repeat(outMessage, repeatCount)
				i++ // Если справа число, то перескакиваем через него.
			}
		}
		outStr.WriteString(outMessage)
	}

	return outStr.String(), nil
}
