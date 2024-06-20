package main

import "fmt"

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 20; i++ {
			ch <- i
		}
	}()
	return ch
}
func main() {
	in := generator()

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)

	for i := 0; i < 20; i++ {
		select {
		case value := <-out1:
			fmt.Println("Output 1 got:", value)
		case value := <-out2:
			fmt.Println("Output 2 got:", value)
		case value := <-out3:
			fmt.Println("Output 3 got:", value)
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for data := range in {
			out <- data
		}
	}()

	return out
}
