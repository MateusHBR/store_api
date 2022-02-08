package product

type service interface {
	getAllProducts() (ProductList, error)
	addProduct(p Product) (Product, error)
}

type productService struct {
	repository repository
}

func (ps productService) getAllProducts() (ProductList, error) {
	productList, err := ps.repository.getAllProducts()

	if err != nil {
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
