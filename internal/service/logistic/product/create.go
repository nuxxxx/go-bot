package product

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

func (ps *DummyProductService) Create(product logistic.Product) (uint64, error) {
	products[product.ID] = product

	return product.ID, nil
}
