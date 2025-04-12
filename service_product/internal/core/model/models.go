package model

type Product struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"imageURL"`
}
