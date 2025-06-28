package product

import (
	"net/http"
	"strconv"
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

	for _, p := range products {
		productResponse := convertToProductResponse(p)
		
		productResponses = append(productResponses, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all products",
		"data": productResponses,
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	p, err := h.productService.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	productResponse := convertToProductResponse(*p)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get product by ID",
		"data": productResponse,
	})
}

func (h *ProductHandler) PostProduct(c *gin.Context) {
	var productRequest CreeateProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	product, err := h.productService.AddNewProduct(productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create new product:" + product.Name,
		"data": convertToProductResponse(*product),
	})
}

// PRODUCT RESPONSE
func convertToProductResponse(p Product) GetProductResponse {
	return GetProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Slug:        p.Slug,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Category:    p.Category,
		Images:      p.Images,
	}
}