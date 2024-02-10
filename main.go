// main.go
package main

import (
	"fmt"
	"sync"
	"time"

	"data_gen/src_csv"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	gofakeit.Seed(time.Now().UnixNano())

	numRows := 100000
	outputFilename := "ecommerce_data.csv"

	ch := make(chan src_csv.Row, 1000)

	var wg sync.WaitGroup

	wg.Add(1)
	go src_csv.GenerateData(numRows, &wg, ch)

	wg.Add(1)
	go src_csv.WriteToCSV(outputFilename, ch, &wg)

	wg.Wait()

	fmt.Printf("Generated %d rows of e-commerce data and saved to %s\n", numRows, outputFilename)
}
