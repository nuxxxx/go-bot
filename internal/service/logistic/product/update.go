package product

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

func (ps *DummyProductService) Update(productID uint64, product logistic.Product) error {
	products[productID] = product

	return nil
}
