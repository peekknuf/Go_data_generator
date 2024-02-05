// src_csv/write.go
package src_csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// WriteToCSV writes data from the provided channel to a CSV file.
func WriteToCSV(filename string, ch <-chan Row, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"id", "timestamp", "product_name", "company", "price", "quantity", "discount", "total_price",
		"customer_id", "first_name", "last_name", "email", "address", "city", "state", "zip", "country",
	}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for row := range ch {
		record := []string{
			strconv.Itoa(row.ID),
			row.Timestamp.Format(time.RFC3339),
			row.ProductName,
			row.Company,
			fmt.Sprintf("%.2f", row.Price),
			strconv.Itoa(row.Quantity),
			fmt.Sprintf("%.2f", row.Discount),
			fmt.Sprintf("%.2f", row.TotalPrice),
			strconv.Itoa(row.CustomerID),
			row.FirstName,
			row.LastName,
			row.Email,
			row.Address,
			row.City,
			row.State,
			row.Zip,
			row.Country,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}
}
