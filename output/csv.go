package output

import (
	"encoding/csv"
	"fmt"
	"os"

	"pricebite/backend/internal/models"
)


func WriteProductsCSV(filename string, products []models.Product) error {
	
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	
	header := []string{"ID", "Name", "PriceValue", "PriceDisplay", "Country"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	
	for _, p := range products {
		row := []string{
			p.ID,
			p.Name,
			fmt.Sprintf("%.2f", p.PriceValue),
			p.PriceDisplay,
			p.Country,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %v", err)
		}
	}

	return nil
}

