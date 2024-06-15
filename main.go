package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 2
	}()
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
