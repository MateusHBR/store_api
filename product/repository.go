package product

import (
	"database/sql"
	"fmt"
)

type repository interface {
	getAllProducts() (ProductList, error)
}

type productRepository struct {
	*sql.DB
}

func (pr productRepository) getAllProducts() (ProductList, error) {
	productList := ProductList{}

	rows, err := pr.Query("SELECT id, name, description, created_at FROM product ORDER BY id DESC")

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
