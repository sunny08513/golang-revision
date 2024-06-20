package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwmu sync.RWMutex

	// Read lock A
	go func() {
		rwmu.RLock()
		fmt.Println("Read lock A acquired")
		time.Sleep(2 * time.Second)
		rwmu.RUnlock()
		fmt.Println("Read lock A released")
	}()

	// Read lock B
	go func() {
		rwmu.RLock()
		fmt.Println("Read lock B acquired")
		time.Sleep(3 * time.Second)
		rwmu.RUnlock()
		fmt.Println("Read lock B released")
	}()

	// Give some time for A and B to acquire the lock
	time.Sleep(100 * time.Millisecond)

	// Write lock C
	go func() {
		rwmu.Lock()
		fmt.Println("Write lock C acquired")
		time.Sleep(1 * time.Second)
		rwmu.Unlock()
		fmt.Println("Write lock C released")
	}()

	// Read lock D
	go func() {
		time.Sleep(200 * time.Millisecond)
		rwmu.RLock()
		fmt.Println("Read lock D acquired")
		rwmu.RUnlock()
		fmt.Println("Read lock D released")
	}()

	// Wait for all goroutines to complete
	time.Sleep(5 * time.Second)
}
