package src_csv

import (
	"fmt"
	"sync"
	"time"

	gf "github.com/brianvoe/gofakeit/v6"
)

// Row represents a data row.
type Row struct {
	ID          int
	Timestamp   time.Time
	ProductName string
	Company     string
	Price       float64
	Quantity    int
	Discount    float64
	TotalPrice  float64
	CustomerID  int
	FirstName   string
	LastName    string
	Email       string
	Address     string
	City        string
	State       string
	Zip         string
	Country     string
}

// GenerateData generates data and sends it to the provided channel.
func GenerateData(numRows int, wg *sync.WaitGroup, ch chan<- Row) {
	defer wg.Done()

	startTime := time.Now()

	for i := 0; i < numRows; i++ {
		price := gf.Price(4.99, 399.99)
		discount := gf.Float64Range(0.0, 0.5)
		startTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		endTime := time.Now()
		row := Row{
			ID:          i + 1,
			Timestamp:   gf.DateRange(startTime, endTime),
			ProductName: gf.CarModel(),
			Company:     gf.Company(),
			Price:       gf.Price(4.99, 399.99),
			Quantity:    gf.Number(1, 50),
			Discount:    discount,
			TotalPrice:  price * (1 - discount),
			CustomerID:  gf.Number(1, 99999),
			FirstName:   gf.FirstName(),
			LastName:    gf.LastName(),
			Email:       gf.Email(),
			Address:     gf.Address().Address,
			City:        gf.City(),
			State:       gf.State(),
			Zip:         gf.Zip(),
			Country:     gf.Country(),
		}

		ch <- row
	}

	close(ch)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Data generation took %s\n", elapsedTime)
}