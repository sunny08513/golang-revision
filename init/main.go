package main

import (
	"fmt"
)

// init function in the first file
func init() {
	fmt.Println("Init function in first file")
}

// Another init function in the same package but different file
func init() {
	fmt.Println("Another init function in the same package")
}

func main() {
	fmt.Println("Main function")
}
