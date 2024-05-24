package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)

	t := TokenCache{
		token:      randomString(15),
		expiration: time.Now().Add(5 * time.Second),
		lock:       &sync.Mutex{},
	}

	go t.start(done)
	for i := 1; i <= 3; i++ {
		fmt.Println(t.GetToken())
		time.Sleep(3 * time.Second)
	}

	close(done)
	fmt.Println("shutdown")
}

type TokenCache struct {
	token      string
	expiration time.Time
	lock       *sync.Mutex
}

func (t *TokenCache) start(done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-time.After(t.expiration.Sub(time.Now())):
			t.lock.Lock()
			fmt.Println("getting new token")
			t.token = randomString(15)
			t.expiration = time.Now().Add(5 * time.Second)
			t.lock.Unlock()
		}
	}
}

func (t *TokenCache) GetToken() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.token
}

func randomString(length int) string {
	//rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
