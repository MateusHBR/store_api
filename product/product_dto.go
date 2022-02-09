package product

import "github.com/MateusHBR/mallus_api/server"

type ProductDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (p ProductDTO) toEntity() Product {
	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}

func (p ProductDTO) fromEntity(entity Product) ProductDTO {
	p.ID = entity.ID
	p.Name = entity.Name
	p.Description = entity.Description
	if !entity.CreatedAt.IsZero() {
		p.CreatedAt = server.FormatTime(entity.CreatedAt)
	}
	if !entity.UpdatedAt.IsZero() {
		p.UpdatedAt = server.FormatTime(entity.UpdatedAt)
	}

	return p
}

type ProductListDTO struct {
	Products []ProductDTO `json:"products"`
}

func (p ProductListDTO) toEntity() ProductList {
	productListEntity := ProductList{}

	for _, productDTO := range p.Products {
		productListEntity.Products = append(productListEntity.Products, productDTO.toEntity())
	}

	return productListEntity
}

func (p ProductListDTO) fromEntity(productList ProductList) ProductListDTO {
	if productList.Products == nil {
		p.Products = make([]ProductDTO, 0)
		return p
	}

	for _, productEntity := range productList.Products {
		p.Products = append(p.Products, ProductDTO{}.fromEntity(productEntity))
	}

	return p
}
