package repository

import (
	"database/sql"
	"fmt"

	"github.com/MateusHBR/mallus_api/product/entity"
)

type ProductRepository struct {
	*sql.DB
}

func (pr ProductRepository) GetAllProducts() (entity.ProductList, error) {
	productList := entity.ProductList{}

	rows, err := pr.Query("SELECT id, name, description, created_at FROM product ORDER BY id DESC")

	if err != nil {
		fmt.Printf("Failed to get query %s \n", err)
		return productList, err
	}

	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)

		if err != nil {
			fmt.Println("Failed to map product")
			return productList, err
		}

		productList.Products = append(productList.Products, p)

	}

	return productList, nil
}
