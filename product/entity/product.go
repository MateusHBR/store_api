package product

type Product struct {
	ID          int
	Name        string
	Description string
	CreatedAt   string
}

type ProductList struct {
	Products []Product
}
