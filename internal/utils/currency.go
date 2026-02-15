package utils

import (
	"strconv"
	"strings"
)


func ParsePriceString(priceStr string) float64 {
	
	clean := strings.ReplaceAll(priceStr, ",", "")
	clean = strings.TrimFunc(clean, func(r rune) bool {
		return r < '0' || r > '9' && r != '.'
	})

	val, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0
	}
	return val
}

