package main

import (
	"fmt"
	"sync"
)

func main() {
	logChannel := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(logChannel chan<- string) {
		defer wg.Done()
		for i := 1; i <= 50; i++ {
			log := fmt.Sprintf("Log Entry %v", i)
			logChannel <- log
		}
		close(logChannel)
	}(logChannel)

	//log processor

	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go func(id int, logChannel <-chan string) {
			defer wg.Done()
			for log := range logChannel {
				fmt.Printf("Processor %v:%v\n", id, log)
			}
		}(id, logChannel)
	}

	wg.Wait()
}
