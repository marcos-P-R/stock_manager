package product

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

func (p *ProductUseCase) GetProductsByCategory(category string) ([]*model.Product, error) {
	products, err := p.productRepository.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	return products, nil
}
