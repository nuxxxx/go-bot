package commands

import (
	"errors"
	"github/nuxxxcake/go-bot/internal/service/logistic/product"
	"strconv"
)

func getProductById(p *product.ProductService, productID string) (string, error) {
	convertedID, err := strconv.Atoi(productID)

	if err != nil {
		return "", errors.New("cannot parse ID")
	}

	if convertedID < 0 {
		return "", errors.New("negative ID")
	}

	found, err := (*p).Describe(uint64(convertedID))

	if err != nil {
		return "", err
	}

	return found.String(), nil
}
