package hw02unpackstring

import (

	// "fmt".
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	runes := []rune(str)
	if len(runes) == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var outStr strings.Builder
	runesLen := len(runes)

	for i := 0; i < runesLen; i++ {
		r := runes[i]
		outMessage := string(r)
		slog.Debug(i, " ", outMessage)

		if unicode.IsDigit(r) {
			// slog.Debug(r, rightRune)
			return "", ErrInvalidString
		}

		if i < runesLen-1 {
			rightRune := runes[i+1]

			if unicode.IsDigit(rightRune) {
				repeatCount, err := strconv.Atoi(string(rightRune))
				if err != nil {
					slog.Panic("что-то неправильно написано....")
				}
				outMessage = strings.Repeat(outMessage, repeatCount)
				i++ // Если справа число, то перескакиваем через него.
			}
		}
		// } else {
		// 	if unicode.IsDigit(r) {
		// 		return "", ErrInvalidString
		// 	}
		// }
		outStr.WriteString(outMessage)
	}

	return outStr.String(), nil
}
