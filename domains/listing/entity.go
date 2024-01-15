package listing

import "time"

type PricePerDay struct {
	Date  time.Time `json:"date"`
	Price float64   `json:"price"`
}