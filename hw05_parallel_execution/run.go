package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

func init() {
	slog = logger.NewSugaredLogger()
}

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrWorkersNotFound     = errors.New("нет свободных воркеров")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {
	// Place your code here.

	if workersCount <= 0 {
		return ErrWorkersNotFound
	}

	var wg sync.WaitGroup
	var errorCount int64 = 0
	var lastTaskPosition int64 = -1
	tasksLen := int64(len(tasks))

	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(tasks []Task, errorCount, lastTaskPosition *int64, maxErrorsCount, runnerID int) {
			defer wg.Done()
			for {
				taskPosition := atomic.AddInt64(lastTaskPosition, 1)
				slog.Debug(fmt.Sprintf("Runner %d, task %d\n", runnerID, taskPosition))
				if taskPosition >= tasksLen {
					return
				}
				if maxErrorsCount > 0 && atomic.LoadInt64(errorCount) >= int64(maxErrorsCount) {
					return
				}
				err := tasks[taskPosition]()
				if err != nil {
					atomic.AddInt64(errorCount, 1)
				}
			}
		}(tasks, &errorCount, &lastTaskPosition, maxErrorsCount, i)
	}

	wg.Wait()
	if maxErrorsCount > 0 && errorCount >= int64(maxErrorsCount) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
