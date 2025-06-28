package product

type CreateProductResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Slug		string	 `json:"slug"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Category    string   `json:"category"`
	Images      string   `json:"images"`
}

type GetProductResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Slug		string	 `json:"slug"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Category    string   `json:"category"`
	Images      string   `json:"images"`
}