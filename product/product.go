package product

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	RemovedAt   time.Time
}

type ProductList struct {
	Products []Product
}
