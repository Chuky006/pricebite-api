package service

import (
	"pricebite/backend/internal/apple"
	"pricebite/backend/internal/models"
)


func FetchAllCountries(countries []string) (map[string][]models.Product, error) {
	result := make(map[string][]models.Product)

	for _, country := range countries {
		products, err := apple.FetchAllIPhones(country)
		if err != nil {
			
			result[country] = []models.Product{}
			continue
		}
		result[country] = products
	}

	return result, nil
}

