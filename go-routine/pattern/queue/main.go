package main

import (
	"fmt"
	"sync"
	"time"
)

const limit = 2
const work = 10

func main() {
	queue := make(chan struct{}, 10)

	wg := sync.WaitGroup{}
	wg.Add(work)

	for i := 0; i < work; i++ {
		process(i, queue, &wg)
	}
	wg.Wait()
}

func process(work int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Processed:", work)

		<-queue
	}()
}

// iptables -t nat -A KUBE-SERVICES -d 10.96.0.1/32 -p tcp --dport 80 -j KUBE-SVC-XXXXX
// iptables -t nat -A KUBE-SVC-XXXXX -m statistic --mode random --probability 0.5 -j KUBE-SEP-YYYYY
// iptables -t nat -A KUBE-SVC-XXXXX -j KUBE-SEP-ZZZZZ
// iptables -t nat -A KUBE-SEP-YYYYY -p tcp -j DNAT --to-destination 192.168.1.1:80
// iptables -t nat -A KUBE-SEP-ZZZZZ -p tcp -j DNAT --to-destination 192.168.1.2:80

/*What is go routine?

- Go routine is execution unit like a light weight thread in another programming language.

What do you mean by light weight thread?
How it differ from normal thread?

How go scheduler works?


*/
