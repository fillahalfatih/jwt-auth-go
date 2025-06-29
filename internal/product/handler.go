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

func (h *ProductHandler) GetAllProductsHandler(c *gin.Context) {
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

func (h *ProductHandler) GetProductByIDHandler(c *gin.Context) {
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
		"message": "Success get product by ID:" + strconv.Itoa(productID),
		"data": productResponse,
	})
}

func (h *ProductHandler) CreateProductHandler(c *gin.Context) {
	var productRequest CreateProductRequest

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

func (h *ProductHandler) UpdateProductHandler(c *gin.Context) {
	var productRequest UpdateProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	product, err := h.productService.UpdateProduct(productID, productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update product with ID:" + strconv.Itoa(productID),
		"data": convertToProductResponse(*product),
	})
}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	product, err := h.productService.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete product with ID:" + strconv.Itoa(productID),
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