package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	mu := sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			count++
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(count)
}
