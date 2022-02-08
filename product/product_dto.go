package product

type ProductDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func (p ProductDTO) toEntity() Product {
	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
	}
}

func (p ProductDTO) fromEntity(entity Product) ProductDTO {
	p.ID = entity.ID
	p.Name = entity.Name
	p.Description = entity.Description
	p.CreatedAt = entity.CreatedAt

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
	for _, productEntity := range productList.Products {
		p.Products = append(p.Products, ProductDTO{}.fromEntity(productEntity))
	}

	return p
}
