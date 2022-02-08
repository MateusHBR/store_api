package product

import (
	"fmt"

	"github.com/MateusHBR/mallus_api/adapter/database"

	product "github.com/MateusHBR/mallus_api/product/entity"
)

type ProductRepository struct {
	database.DatabaseAdapter
}

func (pr ProductRepository) GetAllProducts() (product.ProductList, error) {
	productList := product.ProductList{}
	db, err := pr.OpenConnection()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return productList, err
	}

	rows, err := db.Query("SELECT id, name, description, created_at FROM product ORDER BY id DESC")

	if err != nil {
		fmt.Printf("Failed to get query %s \n", err)
		return productList, err
	}

	for rows.Next() {
		var p product.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)

		if err != nil {
			fmt.Println("Failed to map product")
			return productList, err
		}

		productList.Products = append(productList.Products, p)

	}

	return productList, nil
}
