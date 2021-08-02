package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.

	maxGoroutines := make(chan struct{}, n)

	var wg sync.WaitGroup
	var errorCount int64 = 0

	for _, task := range tasks {
		wg.Add(1)
		go func(task Task, errorCount *int64, m int) {
			defer wg.Done()
			maxGoroutines <- struct{}{}
			if int64(m) > atomic.LoadInt64(errorCount) {
				err := task()
				if err != nil {
					atomic.AddInt64(errorCount, 1)
				}
			}
			<-maxGoroutines
		}(task, &errorCount, m)
	}

	wg.Wait()
	if errorCount >= int64(n) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
