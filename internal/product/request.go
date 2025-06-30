package product

type CreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Slug		string	 `json:"slug" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price      	float64  `json:"price" binding:"required"`
	Quantity    int      `json:"quantity" binding:"required"`
	CategoryID    uint   `json:"category_id" binding:"required"`
	Images      string   `json:"images" binding:"required"`
}

type UpdateProductRequest struct {
	Name        *string   `json:"name"`
	Slug		*string	  `json:"slug"`
	Description *string   `json:"description"`
	Price      	*float64  `json:"price"`
	Quantity    *int      `json:"quantity"`
	CategoryID    *uint   `json:"category_id"`
	Images      *string   `json:"images"`
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

/*
Contoh data JSON untuk POST berdasarkan CreeateProductRequest dengan kategori "Bread":

{
	"name": "Sourdough Bread",
	"slug": "sourdough-bread",
	"description": "Artisan sourdough bread with a crispy crust and tangy flavor.",
	"price": 3.49,
	"quantity": 50,
	"category": "Bread",
	"images": "https://example.com/images/sourdough-bread.jpg"
}
*/