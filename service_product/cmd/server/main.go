package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcos-P-R/stock_manager/service_product/internal/handler"
	"github.com/marcos-P-R/stock_manager/service_product/internal/repository"
	"github.com/marcos-P-R/stock_manager/service_product/internal/usecases/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./products.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		sku TEXT NOT NULL UNIQUE,
		price REAL NOT NULL,
		brand TEXT,
		description TEXT,
		category TEXT,
		image_url TEXT
		);
	`)

	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	repo := repository.NewMySqlRepository(db)
	productUseCase := product.NewProductUseCase(repo)
	handler := handler.NewProductHandler(productUseCase)

	r := gin.Default()
	productGroup := r.Group("/products")
	handler.RegisterRoutes(productGroup)

	r.Run(":8080")
	log.Println("Server running on port 8080")
}
