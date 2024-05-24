package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d processing Job %d\n", id, job)
	}
}

func main() {
	workers := 3
	jobs := 10
	wg := sync.WaitGroup{}

	jobCh := make(chan int)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobCh, &wg)
	}

	for i := 1; i <= jobs; i++ {
		jobCh <- i
	}
	close(jobCh)
	wg.Wait()
}
