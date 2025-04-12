package usecases

import (
	"github.com/marcos-P-R/stock_manager/service_product/internal/core/model"
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

func (p *ProductUseCase) Create(product *model.Product) (bool, error) {
	created, error := p.productRepository.CreateProduct(product)
	if error != nil {
		return false, error
	}
	return created, nil
}
