package product

import (
    "gorm.io/gorm"
)

type Repository interface {
    FindAll() ([]Product, error)
	FindByID(id uint) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(product Product) (Product, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, err
}

func (r *repository) FindByID(id uint) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) CreateProduct(product Product) (Product, error) {
    err := r.db.Create(&product).Error
	if err != nil {
		return Product{}, err
	}

	err = r.db.Preload("Category").First(&product, product.ID).Error
	if err != nil {
		return product, err
	}
	
	return product, nil
}

func (r *repository) UpdateProduct(product Product) (Product, error) {
    err := r.db.Save(&product).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) DeleteProduct(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}