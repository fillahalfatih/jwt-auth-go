package product

type CreeateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Slug		string	`json:"slug" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price      	float64  `json:"price" binding:"required"`
	Quantity   int      `json:"quantity" binding:"required"`
	Category   string   `json:"category" binding:"required"`
	Images     string `json:"images" binding:"required"`
}

/*
Contoh data JSON untuk POST berdasarkan CreeateProductRequest:

{
	"name": "Butter Cookies",
	"slug": "butter-cookies",
	"description": "Delicious homemade butter cookies with a rich, buttery flavor.",
	"price": 5.99,
	"quantity": 100,
	"category": "Cookies",
	"images": "https://example.com/images/butter-cookies.jpg"
}
*/