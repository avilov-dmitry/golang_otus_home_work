package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrInvalidN            = errors.New("n should be positive")
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// M <= 0 means no error limit.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidN
	}

	if m <= 0 {
		m = len(tasks) + 1
	}

	var errCount int32
	taskChan := make(chan Task)
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChan {
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()

	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
