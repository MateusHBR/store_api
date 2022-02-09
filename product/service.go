package product

import (
	"github.com/MateusHBR/mallus_api/adapter/database"
	"github.com/MateusHBR/mallus_api/server"
)

type service interface {
	getAllProducts() (ProductList, error)
	addProduct(p Product) (Product, error)
	updateProduct(p Product) (Product, error)
	deleteProduct(id string) error
}

type productService struct {
	repository repository
}

func (ps productService) getAllProducts() (ProductList, error) {
	productList, err := ps.repository.getAllProducts()

	if err != nil {
		if database.IsNoRowsError(err) {
			return ProductList{}, nil
		}

		return ProductList{}, err
	}

	return productList, nil
}

func (ps productService) addProduct(p Product) (Product, error) {
	product, err := ps.repository.addProduct(p)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (ps productService) updateProduct(p Product) (Product, error) {
	product, err := ps.repository.updateProduct(p)

	if err != nil {
		if database.IsNoRowsError(err) {
			return Product{}, server.ErrorNotFound
		}

		return Product{}, err
	}

	return product, nil
}

func (ps productService) deleteProduct(id string) error {
	err := ps.repository.deleteProduct(id)

	if err != nil {
		if database.IsNoRowsError(err) {
			return server.ErrorNotFound
		}

		return err
	}

	return nil
}
