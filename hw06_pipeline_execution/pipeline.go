package hw06pipelineexecution

import (
	"sync"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

func init() {
	slog = logger.NewSugaredLogger()
}

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func execStages(data interface{}, stages ...Stage) Out {
	out := make(Bi)

	go func() {
		fIn := make(Bi)
		if len(stages) > 0 {
			fOut := stages[0](fIn)
			for i := 1; i < len(stages); i++ {
				fOut = stages[i](fOut)
			}
			fIn <- data
			out <- <-fOut
		} else {
			out <- data
		}

		close(out)
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var mapMutex = &sync.Mutex{}
	var doneMutex = &sync.Mutex{}

	out := make(Bi)
	closed := false

	go func() {
		<-done
		slog.Debug("донон")
		doneMutex.Lock()
		slog.Debug("лок донона")
		if !closed {
			slog.Debug("закрытие")
			close(out)
			closed = true
		}
		doneMutex.Unlock()
	}()

	go func() {
		// values := make(map[int]Out)
		values := make(map[int]interface{})
		var wg sync.WaitGroup
		i := -1
		for data := range in {
			wg.Add(1)
			i++
			slog.Debug("Data: ", data)
			go func(values map[int]interface{}, index int, data interface{}, stages ...Stage) {
				defer wg.Done()
				value := execStages(data, stages...)
				mapMutex.Lock()
				values[index] = <-value
				mapMutex.Unlock()
			}(values, i, data, stages...)
		}
		wg.Wait()

		doneMutex.Lock()
		slog.Debug("closed: ", closed)
		if !closed {
			slog.Debug("вывод")

			for j := 0; j <= i; j++ {
				v := values[j]
				out <- v
			}
			close(out)
			closed = true
		}
		doneMutex.Unlock()
	}()
	return out
}
