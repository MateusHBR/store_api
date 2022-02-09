package product

import (
	"database/sql"
	"fmt"

	"github.com/MateusHBR/mallus_api/adapter/database"
	"github.com/MateusHBR/mallus_api/server"
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

	sqlStmt := `
	SELECT id, name, description, created_at, updated_at
	FROM products`
	rows, err := pr.Query(sqlStmt)

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
	p.CreatedAt = server.TimeNow()
	p.UpdatedAt = p.CreatedAt

	sqlStmt := `
	INSERT INTO products(id, name, description, created_at, updated_at)
	values($1, $2, $3, $4, $5)`
	_, err := pr.Exec(sqlStmt, p.ID, p.Name, p.Description, p.CreatedAt, p.UpdatedAt)
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
	p.UpdatedAt = server.TimeNow()

	sqlStmt := `
	UPDATE products
	SET name=$1, description=$2, updated_at=$3
	WHERE id=$4`
	res, err := pr.Exec(sqlStmt, p.Name, p.Description, p.UpdatedAt, p.ID)

	if err != nil {
		return Product{}, err
	}

	if err := database.CheckHasNoRowsAffected(res); err != nil {
		return Product{}, err
	}

	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}

func (pr productRepository) deleteProduct(id string) error {

	res, err := pr.Exec("DELETE FROM products WHERE id=$1", id)

	if err != nil {
		return err
	}

	return database.CheckHasNoRowsAffected(res)
}
