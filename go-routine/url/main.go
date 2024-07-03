package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	ID  int
	URL string
}

func fetchurl(url string, id int, cancelCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	resp := &http.Response{}

	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		resp, _ = http.Get(url)
		close(done)
	}()

	select {
	case <-cancelCh:
		fmt.Printf("Task %d: Request to %s was cancelled\n", id, url)
		return
	case <-done:
		fmt.Printf("Task %d: Response from %s: %s\n", id, url, resp.Status)
		return
	}
}

func main() {
	urls := []Task{
		{ID: 1, URL: "http://example.com"},
		{ID: 2, URL: "http://example.org"},
		{ID: 3, URL: "http://example.net"},
	}

	cancelChannels := make(map[int]chan struct{})
	var mu sync.Mutex

	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)

		cancelChannel := make(chan struct{})
		mu.Lock()
		cancelChannels[url.ID] = cancelChannel
		mu.Unlock()
		go fetchurl(url.URL, url.ID, cancelChannel, wg)
	}

	mu.Lock()
	if cancelChan, exists := cancelChannels[2]; exists {
		close(cancelChan)
		delete(cancelChannels, 2)
	}
	mu.Unlock()

	// Wait for all remaining goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
