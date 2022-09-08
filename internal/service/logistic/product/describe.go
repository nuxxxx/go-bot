package product

import (
	"errors"
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

func (ps *DummyProductService) Describe(productID uint64) (*logistic.Product, error) {
	product, ok := products[productID]

	if ok {
		return &product, nil
	}

	return nil, errors.New("there is no product with this ID")
}
