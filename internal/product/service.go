package product

type Service interface {
	GetAllProduct() ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	AddNewProduct(product CreateProductRequest) (*Product, error)
	UpdateProduct(ID int, product UpdateProductRequest) (*Product, error)
	DeleteProduct(ID int) (*Product, error)
}

type service struct {
    repository Repository
}

func NewService(repository Repository) Service {
    return &service{repository: repository}
}

func (s *service) GetAllProduct() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
        return nil, err
    }
	return products, err
}

func (s *service) GetProductByID(id uint) (*Product, error) {
    product, err := s.repository.FindByID(id)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (s *service) AddNewProduct(productRequest CreateProductRequest) (*Product, error) {
	newProduct := Product{
		Name:        productRequest.Name,
		Slug:        productRequest.Slug,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Quantity:    productRequest.Quantity,
		Category:    productRequest.Category,
		Images:      productRequest.Images,
	}

	product, err := s.repository.CreateProduct(newProduct)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *service) UpdateProduct(ID int, productRequest UpdateProductRequest) (*Product, error) {
	existingProduct, err := s.repository.FindByID(uint(ID))
	if err != nil {
		return nil, err
	}

	if productRequest.Name != nil {
		existingProduct.Name = *productRequest.Name
	}
	if productRequest.Slug != nil {
		existingProduct.Slug = *productRequest.Slug
	}
	if productRequest.Description != nil {
		existingProduct.Description = *productRequest.Description
	}
	if productRequest.Price != nil {
		existingProduct.Price = *productRequest.Price
	}
	if productRequest.Quantity != nil {
		existingProduct.Quantity = *productRequest.Quantity
	}
	if productRequest.Category != nil {
		existingProduct.Category = *productRequest.Category
	}
	if productRequest.Images != nil {
		existingProduct.Images = *productRequest.Images
	}

	res, err := s.repository.UpdateProduct(existingProduct)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *service) DeleteProduct(ID int) (*Product, error) {
	product, err := s.repository.FindByID(uint(ID))
	if err != nil {
		return nil, err
	}

	deletedProduct, err := s.repository.DeleteProduct(product)
	if err != nil {
		return nil, err
	}

	return &deletedProduct, nil
}