package repository

import (
	"database/sql"

	"github.com/marcos-P-R/stock_manager/service_product/internal/core/model"
	"github.com/marcos-P-R/stock_manager/service_product/internal/core/ports"
)

type mySQLProductRepository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) ports.ProductRepository {
	return &mySQLProductRepository{
		db: db,
	}
}

func (r *mySQLProductRepository) CreateProduct(product *model.Product) (bool, error) {
	query := "INSERT INTO products (name, price, brand, description, category, imageurl) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Price, product.Brand, product.Description, product.Category, product.ImageURL)
	if err != nil {
		return false, err
	}
	return true, nil
}
