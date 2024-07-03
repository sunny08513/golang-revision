package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func readCSV(filepath string, rowCh chan<- map[string]string) {
	defer close(rowCh)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error occurred during file open:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the column names
	columns, err := reader.Read()
	if err != nil {
		fmt.Println("error during read columns data:", err)
		return
	}

	// Read the rows and send to channel
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		// Create a map for the row
		rowData := make(map[string]string)
		for i, col := range columns {
			rowData[col] = row[i]
		}

		rowCh <- rowData
	}
}

func convertToJson(rowCh <-chan map[string]string, jsonCh chan<- []byte) {
	defer close(jsonCh)
	for row := range rowCh {
		jsonData, err := json.MarshalIndent(row, "", "  ")
		if err != nil {
			fmt.Println("error during JSON marshal:", err)
			return
		}

		jsonCh <- jsonData
	}
}

func main() {
	rowCh := make(chan map[string]string)
	jsonCh := make(chan []byte)

	csvFilePath := "second/sample2.csv"
	jsonFilePath := "sample2.json"

	// Read CSV and convert to JSON concurrently
	go readCSV(csvFilePath, rowCh)
	go convertToJson(rowCh, jsonCh)

	// Collect all JSON data
	var jsonData []byte
	jsonData = append(jsonData, '[')
	first := true
	for jsonRow := range jsonCh {
		if !first {
			jsonData = append(jsonData, ',')
		}
		first = false
		jsonData = append(jsonData, jsonRow...)
	}
	jsonData = append(jsonData, ']')

	// Write JSON data to file
	err := os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Println("error occurred during writing JSON file:", err)
		return
	}

	fmt.Println("CSV file has been successfully converted to JSON")
}
