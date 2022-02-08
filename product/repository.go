package product

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type repository interface {
	getAllProducts() (ProductList, error)
	addProduct(p Product) (Product, error)
}

type productRepository struct {
	*sql.DB
}

func (pr productRepository) getAllProducts() (ProductList, error) {
	productList := ProductList{}

	rows, err := pr.Query("SELECT id, name, description, created_at FROM product")

	if err != nil {
		fmt.Printf("Failed to get query %s \n", err)
		return productList, err
	}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)

		if err != nil {
			fmt.Println("Failed to map product")
			return productList, err
		}

		productList.Products = append(productList.Products, p)

	}

	return productList, nil
}

func (pr productRepository) addProduct(p Product) (Product, error) {
	p.ID = uuid.New().String()
	productCreatedAt := time.Now()
	p.CreatedAt = productCreatedAt.Format(time.RFC3339)

	_, err := pr.Exec("INSERT INTO product(id, name, description, created_at) values($1, $2, $3, $4)", p.ID, p.Name, p.Description, productCreatedAt)
	if err != nil {
		fmt.Printf("Failed to insert product %s \n", err)
		return Product{}, err
	}

	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
	}, nil
}
