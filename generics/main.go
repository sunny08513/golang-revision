package main

import "fmt"

// Generic struct
type Pair[K, V any] struct {
	Key   K
	Value V
}

// Generic interface
type Container[T any] interface {
	Add(item T)
	Remove()
}

// custom constraint using an interface
type sumInterface interface {
	int | float64 | string
}

// Generic function
func sum[T sumInterface](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.5, 2.0))
	fmt.Println(sum("a", "b"))
}
