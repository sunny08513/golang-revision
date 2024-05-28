package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	ID    int
	Name  string
	Price string
}

func main() {
	filePath := "file/csv/product.csv"

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

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error in readin %v", records)
		return
	}

	var products []Product
	for idx, record := range records[1:] { // Skip the header row
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Error to read id for row number: %v\n", idx)
			continue
		}
		product := Product{
			ID:    id,
			Name:  record[1],
			Price: record[2],
		}
		products = append(products, product)
	}

	fmt.Println(products)
}
