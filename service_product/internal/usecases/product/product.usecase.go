package product

import (
	"github.com/marcos-P-R/stock_manager/service_product/internal/core/ports"
)

type ProductUseCase struct {
	productRepository ports.ProductRepository
}

func NewProductUseCase(productRepository ports.ProductRepository) ports.ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}
