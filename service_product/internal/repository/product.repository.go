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
	query := "INSERT INTO products (name, sku, price, brand, description, category, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Sku, product.Price, product.Brand, product.Description, product.Category, product.ImageURL)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mySQLProductRepository) GetProductBySKU(sku string) (*model.Product, error) {
	query := "SELECT name, sku, price, brand, description, category, image_url FROM products WHERE sku = ?"
	row := r.db.QueryRow(query, sku)
	product := &model.Product{}

	err := row.Scan(&product.Name, &product.Sku, &product.Price, &product.Brand, &product.Description, &product.Category, &product.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

func (r *mySQLProductRepository) GetProductsByCategory(category string) ([]*model.Product, error) {
	query := "SELECT name, sku, price, brand, description, category, image_url FROM products WHERE category = ?"
	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*model.Product{}

	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.Name, &product.Sku, &product.Price, &product.Brand, &product.Description, &product.Category, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
