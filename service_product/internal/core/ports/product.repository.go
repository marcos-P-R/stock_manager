package ports

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

type ProductRepository interface {
	CreateProduct(*model.Product) (bool, error)
	GetProductBySKU(sku string) (*model.Product, error)
}
