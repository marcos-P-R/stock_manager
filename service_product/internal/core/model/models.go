package model

type Product struct {
	Name        string  `json:"name"`
	Sku         string  `json:"sku"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"imageURL"`
}
