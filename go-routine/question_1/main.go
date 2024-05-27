package main

import (
	"fmt"
	"sync"
)

func main() {
	numChannel := make(chan int, 10)
	squareChannel := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(numChannel chan<- int) {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			numChannel <- i
		}
		close(numChannel)
	}(numChannel)

	go func(numChannel <-chan int, squareChannel chan<- int) {
		defer wg.Done()
		for val := range numChannel {
			squareChannel <- val * val
		}
		close(squareChannel)
	}(numChannel, squareChannel)

	go func(squareChannel <-chan int) {
		defer wg.Done()
		for val := range squareChannel {
			fmt.Println(val)
		}
	}(squareChannel)

	wg.Wait()
}
