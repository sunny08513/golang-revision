package main

import "fmt"

// Define a function that returns another function
func getMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// Define a function that takes an int and returns an int
func square(n int) int {
	return n * n
}

// Define a function that takes a function as an argument
func applyFunction(n int, f func(int) int) int {
	return f(n)
}

func main() {
	// Get a multiplier function
	double := getMultiplier(2)
	triple := getMultiplier(3)

	// Use the returned functions
	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15

	// Pass the function as an argument
	result := applyFunction(5, square)
	fmt.Println(result) // Output: 25
}
