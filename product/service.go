package product

type service interface {
	getAllProducts() (ProductList, error)
}

type productService struct {
	repository repository
}

func (pr productService) getAllProducts() (ProductList, error) {
	productList, err := pr.repository.getAllProducts()

	if err != nil {
		return ProductList{}, err
	}

	return productList, nil
}
