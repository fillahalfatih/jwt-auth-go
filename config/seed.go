package config

import (
	"jwt-auth-go/internal/product"
	"log"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categories := []product.Category{
		{
			Name:        "Donut",
		},
		{
			Name:        "Cookies",
		},
	}

	for _, c := range categories {
		err := db.Create(&c).Error
		if err != nil {
			log.Printf("Gagal seed kategori: %v\n", err)
		}
	}
}

func SeedProducts(db *gorm.DB) {
	products := []product.Product{
		{
			Name:        "Butter Cookies Matcha",
			Slug:        "butter-cookies-matcha",
			Description: "Delicious cookies with matcha flavour.",
			Price:       1.59,
			Quantity:    21,
			CategoryID:  2,
			Images:      "https://example.com/images/matcha-cookies.jpg",
		},
		{
			Name:        "Donut Cokelat",
			Slug:        "donut-cokelat",
			Description: "Donut dengan taburan cokelat meleleh.",
			Price:       1.25,
			Quantity:    30,
			CategoryID:  1,
			Images:      "https://example.com/images/donut-cokelat.jpg",
		},
	}

	for _, p := range products {
		err := db.Create(&p).Error
		if err != nil {
			log.Printf("Gagal seed produk: %v\n", err)
		}
	}
}
