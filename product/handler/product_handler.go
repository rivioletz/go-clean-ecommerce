package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rivioletz/go-clean-ecommerce/product/dto"
	"github.com/rivioletz/go-clean-ecommerce/product/usecase"
)

type ProductHandler struct {
	UseCase usecase.ProductUseCase
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(uc usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		UseCase: uc,
	}
}

// CreateProduct handles the creation of a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	product, err := h.UseCase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct handles fetching a product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.UseCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
