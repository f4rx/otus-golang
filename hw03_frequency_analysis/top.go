package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var validWord = regexp.MustCompile(`^[^-][\p{L}-]*`)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

func init() {
	slog = logger.NewSugaredLogger()
}

func Top10(str string) []string {
	// slog.Debug(str)
	words := strings.Fields(str)
	slog.Debug(words)
	m := generateMap(words)
	w := sortWordKyes(m)
	if len(w) > 10 {
		return w[:10]
	} else { //nolint:golint
		return w // мне нравится более явное обозначание в таких конструкциях
	}
}

func generateMap(words []string) map[string]int64 {
	wordsCounted := make(map[string]int64)
	for _, word := range words {
		lowWord := strings.ToLower(word)
		strippedWord := validWord.FindString(lowWord)
		if len(strippedWord) > 0 {
			wordsCounted[strippedWord]++
		}
	}
	slog.Debug(wordsCounted)
	return wordsCounted
}

func sortWordKyes(wordsMap map[string]int64) []string {
	keys := make([]string, 0, len(wordsMap))
	for key := range wordsMap {
		keys = append(keys, key)
	}

	sortFunc := func(i, j int) bool {
		if wordsMap[keys[i]] > wordsMap[keys[j]] {
			return true
		} else if wordsMap[keys[i]] == wordsMap[keys[j]] {
			return keys[i] < keys[j]
		}
		return false
	}

	sort.SliceStable(keys, sortFunc)

	slog.Debug(keys)
	return keys
}
