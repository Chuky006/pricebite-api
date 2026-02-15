package apple

import (
	"fmt"
	"log"
	"strings"

	"pricebite/backend/internal/models"
	"pricebite/backend/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

var iPhoneModels = []string{
	"iphone-17-pro",
	"iphone-17",
	"iphone-air",
	"iphone-16",
	"iphone-16e",
}

func FetchAllIPhones(country string) ([]models.Product, error) {
	var products []models.Product

	for _, model := range iPhoneModels {
		url := fmt.Sprintf("https://www.apple.com/%s/shop/buy-iphone/%s", country, model)
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Printf("failed to fetch %s: %v", url, err)
			continue
		}

		name := strings.TrimSpace(doc.Find("h1.as-productname").Text())
		if name == "" {
			name = strings.ReplaceAll(strings.Title(strings.ReplaceAll(model, "-", " ")), "Iphone", "iPhone")
		}

		price := strings.TrimSpace(doc.Find("span.as-price-currentprice").First().Text())
		if price == "" {
			price = strings.TrimSpace(doc.Find("span.current_price").First().Text())
		}

		if name != "" && price != "" {
			
			numericPrice := utils.ParsePriceString(price)

			products = append(products, models.Product{
				Name:         name,
				PriceDisplay: price,
				PriceValue:   numericPrice,
				Country:      country,
			})
		}
	}

	return products, nil
}
