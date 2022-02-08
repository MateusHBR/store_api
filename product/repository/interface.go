package product

import (
	entity "github.com/MateusHBR/mallus_api/product/entity"
)

type Repository interface {
	GetAllProducts() (entity.ProductList, error)
}
