package main

import (
	"errors"
	"log"
	"sync"
)

func worker(wg *sync.WaitGroup) error {
	defer wg.Done()
	// Simulating an error
	return errors.New("something went wrong")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Println("Hi")
		if err := worker(&wg); err != nil {
			log.Printf("error occured: %v", err)
		}
	}()
	wg.Wait()
}
