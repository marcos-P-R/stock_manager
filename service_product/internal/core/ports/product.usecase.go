package ports

import "github.com/marcos-P-R/stock_manager/service_product/internal/core/model"

type ProductUseCase interface {
	Create(product *model.Product) (bool, error)
}
