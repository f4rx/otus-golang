package hw05parallelexecution

import (
	"context"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, tasksChan chan Task, errorsChan chan error, workerID int) {
	defer wg.Done()

	for task := range tasksChan {
		select {
		case <-ctx.Done():
			slog.Debug("stop worker ", workerID)
			return
		default:
		}

		err := task()
		errorsChan <- err
	}
}

func errorUpdater(errorsChan chan error, maxErrorsCount int, cancel context.CancelFunc) {
	accumErrors := 0
	isCanceled := false
	for err := range errorsChan {
		if err != nil {
			accumErrors++
		}
		slog.Debug("Errors: ", accumErrors)
		if maxErrorsCount > 0 && accumErrors >= maxErrorsCount {
			if !isCanceled {
				slog.Debug("cancel()")
				cancel()
				isCanceled = true
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run2(tasks []Task, workersCount, maxErrorsCount int) error {
	// Place your code here.

	if workersCount <= 0 {
		return ErrWorkersNotFound
	}

	var wg sync.WaitGroup
	tasksChan := make(chan Task)
	errorsChan := make(chan error)
	wg.Add(workersCount)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < workersCount; i++ {
		i := i
		go worker(ctx, &wg, tasksChan, errorsChan, i)
	}

	go errorUpdater(errorsChan, maxErrorsCount, cancel)

	isErrorLimitExceeded := false //nolint:ifshort
	/*
			run.go:82:2: variable 'isErrorLimitExceeded' is only used in the if-statement (run.go:106:2); consider using short syntax (ifshort)
		        isErrorLimitExceeded := false
				Почему-то линтер хочет:
				var isErrorLimitExceeded bool
				Вроде нифига не понятнее
	*/

FORLOOP:
	for i, task := range tasks {
		select {
		case <-ctx.Done():
			slog.Debug("stop")
			isErrorLimitExceeded = true
			close(tasksChan)
			break FORLOOP
		case tasksChan <- task:
			slog.Debug("send task ", i)
		}
	}

	cancel()
	slog.Debug("wg wait")
	if !isErrorLimitExceeded {
		close(tasksChan)
	}
	wg.Wait()

	close(errorsChan)

	if isErrorLimitExceeded {
		return ErrErrorsLimitExceeded
	}

	return nil
}
