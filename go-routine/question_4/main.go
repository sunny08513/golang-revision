package main

import (
	"fmt"
	"sync"
)

func main() {
	taskGeneratorCh := make(chan string)
	taskDone := make(chan struct{})
	wg := sync.WaitGroup{}
	workerWg := sync.WaitGroup{}

	wg.Add(1)
	go func(taskGeneratorCh chan<- string) {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			log := fmt.Sprintf("Task %v", i)
			taskGeneratorCh <- log
		}
		close(taskGeneratorCh)
	}(taskGeneratorCh)

	//workers go routine

	for id := 1; id <= 5; id++ {
		workerWg.Add(1)
		go func(id int, taskGeneratorCh <-chan string) {
			defer workerWg.Done()
			for task := range taskGeneratorCh {
				fmt.Printf("Process %v by worker %v\n", task, id)
			}
		}(id, taskGeneratorCh)
	}

	go func() {
		workerWg.Wait()
		taskDone <- struct{}{}
		close(taskDone)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-taskDone:
				fmt.Println("All tasks completed")
				return
			default:
				fmt.Println("Task still going on")
			}
		}
	}()

	wg.Wait()
}
