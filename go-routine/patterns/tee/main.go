package main

import (
	"fmt"
	"sync"
)

// Tee function splits a single input channel into multiple output channels.
func Tee(input <-chan int, outputs ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(outputs))

	// Start a goroutine for each output channel
	for _, out := range outputs {
		go func(out chan int) {
			defer wg.Done()
			for val := range input {
				out <- val
			}
			close(out)
		}(out)
	}

	// Wait for all output goroutines to finish
	go func() {
		wg.Wait()
	}()
}

func main() {
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)

	// Create the tee channel
	go Tee(input, output1, output2)

	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Consumer for output1
	go func() {
		for val := range output1 {
			fmt.Printf("Output1: %d\n", val)
		}
	}()

	// Consumer for output2
	go func() {
		for val := range output2 {
			fmt.Printf("Output2: %d\n", val)
		}
	}()

	// Wait for a key press to exit
	fmt.Scanln()
}
