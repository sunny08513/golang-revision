package main

import (
	"fmt"
)

func doSomething() error {
	// Simulating an error
	return fmt.Errorf("an error occurred")
}

func main() {
	// Create a channel to communicate errors
	errCh := make(chan error)

	// Start a goroutine
	go func() {
		// Perform some operation that might return an error
		if err := doSomething(); err != nil {
			// Send the error to the channel
			errCh <- err
		}
	}()

	// Use a select statement to wait for the goroutine to finish
	select {
	case err := <-errCh:
		// Handle the error
		fmt.Printf("Error occurred: %v\n", err)
	}
}
