package hw03frequencyanalysis

import (
	"strings"
	"sort"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

func init() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck
}

/*
regexp.MustCompile
strings.Split
strings.Fields
sort.Slice
*/


func Top10(str string) []string {
	// slog.Debug(str)
	words := strings.Fields(str)
	slog.Debug(words)
	m := generateMap(words)
	w := sortWordKyes(m)
	if len(w) > 10 {
		return w[:10]
	} else {
		return w
	}
}

func generateMap(words []string) map[string]int64 {
	wordsCounted := make(map[string]int64)
	for _, word := range words {
		wordsCounted[word]++
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
