package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-P-R/stock_manager/service_product/internal/core/model"
	"github.com/marcos-P-R/stock_manager/service_product/internal/core/ports"
)

type ProductHandler struct {
	product ports.ProductUseCase
}

func NewProductHandler(product ports.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		product: product,
	}
}

func (h *ProductHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/", h.CreateProduct)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	created, err := h.product.Create(&product)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	if !created {
		c.JSON(500, gin.H{"error": "Failed to create product"})
	}

	c.JSON(201, gin.H{"message": "Product created successfully"})
}
