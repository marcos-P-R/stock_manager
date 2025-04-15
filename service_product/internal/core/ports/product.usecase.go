package ports

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

type ProductUseCase interface {
	Create(product *model.Product) (bool, error)
	GetProductBySKU(sku string) (*model.Product, error)
	GetProductsByCategory(category string) ([]*model.Product, error)
}
