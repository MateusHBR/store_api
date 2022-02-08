package product

type Product struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
}

type ProductList struct {
	Products []Product
}
