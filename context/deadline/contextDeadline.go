package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	newCtx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("close goroutine")
				return
			}
			// do something else
		}
	}(newCtx)

	<-newCtx.Done()
	fmt.Println("finished")
}
