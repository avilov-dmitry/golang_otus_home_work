package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChan := make(chan Task)
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChan {
				task()
			}
		}()
	}

	for _, task := range tasks {
		taskChan <- task
	}

	close(taskChan)

	wg.Wait()
	return nil
}
