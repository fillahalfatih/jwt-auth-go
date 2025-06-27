package product

import (
    "gorm.io/gorm"
)

type Repository interface {
	CreateProduct(product *Product) error
    FindByID(id uint) (*Product, error)
    FindAll() ([]Product, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Implementasi method CreateProduct
func (r *repository) CreateProduct(product *Product) error {
    return r.db.Create(product).Error
}

func (r *repository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err // Akan mengembalikan error, misal gorm.ErrRecordNotFound jika tidak ada
	}
	return &product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err // Akan mengembalikan error, misal gorm.ErrRecordNotFound jika tidak ada
	}
	return products, nil
}