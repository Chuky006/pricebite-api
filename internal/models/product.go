package models

type Product struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	PriceValue   float64 `json:"priceValue"`
	PriceDisplay string  `json:"priceDisplay"`
	Country      string  `json:"country"`
}

