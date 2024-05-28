package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	filePath := "file/json/product.json"
	productsByte, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("file doesn't exist %v", err)
		} else {
			fmt.Printf("Error opening file %v", err)
		}
		return
	}

	products := []Product{}
	if err = json.Unmarshal(productsByte, &products); err != nil {
		fmt.Printf("Unable to marshal %v", err)
		return
	}
	fmt.Println(products)
}
