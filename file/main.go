package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "file/test.txt"

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File does not exist: %s\n", filePath)
		} else {
			fmt.Printf("Error opening file: %v\n", err)
		}
		return
	}
	defer file.Close()

	filePath = "file/test.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File does not exist: %s\n", filePath)
		} else {
			fmt.Printf("Error opening file: %v\n", err)
		}
		return
	}
	fmt.Println(string(data))
}
