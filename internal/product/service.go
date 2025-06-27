package product

type Service interface {
	GetAllProduct() ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	AddNewProduct(product CreeateProductRequest) (*Product, error)
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
    return product, nil
}

func (s *service) AddNewProduct(productRequest CreeateProductRequest) (*Product, error) {
	newProduct := Product{
		Name:        productRequest.Name,
		Slug:        productRequest.Slug,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Quantity:    productRequest.Quantity,
		Category:    productRequest.Category,
		Images:      productRequest.Images,
	}

	err := s.repository.CreateProduct(&newProduct)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}