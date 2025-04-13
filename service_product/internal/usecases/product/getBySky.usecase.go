package product

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

func (p *ProductUseCase) GetProductBySKU(sku string) (*model.Product, error) {
	product, err := p.productRepository.GetProductBySKU(sku)
	if err != nil {
		return nil, err
	}

	return product, nil
}
