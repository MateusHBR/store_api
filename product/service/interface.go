package product

import (
	product "github.com/MateusHBR/mallus_api/product/dto"
)

type Service interface {
	GetAllProducts() (product.ProductListDTO, error)
}
