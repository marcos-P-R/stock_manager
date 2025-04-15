package product

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

func (p *ProductUseCase) Create(product *model.Product) (bool, error) {
	created, error := p.productRepository.CreateProduct(product)
	if error != nil {
		return false, error
	}
	return created, nil
}
