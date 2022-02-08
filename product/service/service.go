package service

import (
	"github.com/MateusHBR/mallus_api/product/dto"
	"github.com/MateusHBR/mallus_api/product/repository"
)

type ProductService struct {
	repository.Repository
}

func (pr ProductService) GetAllProducts() (dto.ProductListDTO, error) {
	productList, err := pr.Repository.GetAllProducts()

	if err != nil {
		return dto.ProductListDTO{}, err
	}

	var productListDTO = dto.ProductListDTO{}

	for _, p := range productList.Products {
		var productDto = dto.ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
		}

		productListDTO.Products = append(productListDTO.Products, productDto)
	}

	return productListDTO, nil

}
