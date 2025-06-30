package product

import (
	"gorm.io/gorm"
)

type Product struct {
	Name     string  `json:"name" gorm:"not null"`
	Slug     string  `json:"slug" gorm:"not null;unique"`
	Description string `json:"description" gorm:"type:text"`
	Price    float64 `json:"price" gorm:"not null"`
	Quantity int     `json:"quantity" gorm:"not null"`
	Images   string `json:"images" gorm:"type:text"`
	CategoryID uint `json:"category_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	gorm.Model
}

type Category struct {
	Name string `json:"name" gorm:"not null;unique"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID;references:ID"`
	gorm.Model
}