package dto

type ProductDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// func (p ProductDTO) toEntity() product.Product {
// 	return product.Product{
// 		ID:          p.ID,
// 		Name:        p.Name,
// 		Description: p.Description,
// 		CreatedAt:   p.CreatedAt,
// 	}
// }

type ProductListDTO struct {
	Products []ProductDTO `json:"products"`
}

// func (p ProductListDTO) toEntity() product.ProductList {
// 	productListEntity := product.ProductList{}

// 	for _, p := range productListEntity.Products {

// 		productListEntity.Products = append(productListEntity.Products, p)
// 	}

// }
