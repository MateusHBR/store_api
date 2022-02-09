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
	updateProduct(p Product) (Product, error)
	deleteProduct(id string) error
}

type productRepository struct {
	*sql.DB
}

func (pr productRepository) getAllProducts() (ProductList, error) {
	productList := ProductList{}

	rows, err := pr.Query("SELECT id, name, description, created_at, updated_at FROM product")

	if err != nil {
		fmt.Printf("Failed to get query %s \n", err)
		return productList, err
	}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt)

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
	p.UpdatedAt = p.CreatedAt

	_, err := pr.Exec("INSERT INTO product(id, name, description, created_at, updated_at) values($1, $2, $3, $4, $5)", p.ID, p.Name, p.Description, productCreatedAt, productCreatedAt)
	if err != nil {
		fmt.Printf("Failed to insert product %s \n", err)
		return Product{}, err
	}

	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}

func (pr productRepository) updateProduct(p Product) (Product, error) {
	productUpdatedAt := time.Now()
	p.UpdatedAt = productUpdatedAt.Format(time.RFC3339)

	_, err := pr.Exec("UPDATE product SET name=$1, description=$2, updated_at=$3 WHERE id=$4", p.Name, p.Description, p.UpdatedAt, p.ID)

	if err != nil {
		fmt.Printf("Failed to insert product %s \n", err)
		return Product{}, err
	}

	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}

func (pr productRepository) deleteProduct(id string) error {

	_, err := pr.Exec("DELETE FROM product WHERE id=$1", id)

	if err != nil {
		fmt.Printf("Failed to insert product %s \n", err)
		return err
	}

	return nil
}
