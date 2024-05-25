package main

import (
	"fmt"
	"sync"
)

func main() {
	str := "Hello"
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(str string) {
		fmt.Println(str + ", goroutine!")
		wg.Done()
	}(str)
	//time.Sleep(time.Second)
	wg.Wait()
}
