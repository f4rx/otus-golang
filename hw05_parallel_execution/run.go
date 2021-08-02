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

	// maxGoroutines := make(chan struct{}, n)

	var wg sync.WaitGroup
	var errorCount int64 = 0
	var lastTaskPosition int64 = -1

	for i := 0; i <= n; i++ {
		wg.Add(1)
		go func(tasks []Task, errorCount, lastTaskPosition *int64, maxErrorsCount, runnerID int) {
			defer wg.Done()
			for {
				taskPosition := atomic.AddInt64(lastTaskPosition, 1)
				// fmt.Printf("Runner %d, task %d\n", runnerID, taskPosition)
				if taskPosition >= int64(len(tasks)) {
					return
				}
				if atomic.LoadInt64(errorCount) >= int64(maxErrorsCount) {
					return
				}
				err := tasks[taskPosition]()
				if err != nil {
					atomic.AddInt64(errorCount, 1)
				}
			}
		}(tasks, &errorCount, &lastTaskPosition, m, i)
	}

	// for _, task := range tasks {
	// 	wg.Add(1)
	// 	go func(task Task, errorCount *int64, m int) {
	// 		defer wg.Done()
	// 		maxGoroutines <- struct{}{}
	// 		if int64(m) > atomic.LoadInt64(errorCount) {
	// 			err := task()
	// 			if err != nil {
	// 				atomic.AddInt64(errorCount, 1)
	// 			}
	// 		}
	// 		<-maxGoroutines
	// 	}(task, &errorCount, m)
	// }

	wg.Wait()
	if errorCount >= int64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
