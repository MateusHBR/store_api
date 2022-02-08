package product

type Product struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	RemovedAt   string
}

type ProductList struct {
	Products []Product
}
