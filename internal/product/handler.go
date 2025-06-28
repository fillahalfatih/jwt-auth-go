package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService   Service
}

func NewProductHandler(productService Service) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productResponses []GetProductResponse
	for _, product := range products {
		productResponses = append(productResponses, GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Category:    product.Category,
			Images:      product.Images,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all products",
		"products": productResponses,
	})
}