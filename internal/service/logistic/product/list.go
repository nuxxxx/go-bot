package product

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

const LIMIT = 10

func (ps *DummyProductService) List(cursor uint64, limit uint64) ([]logistic.Product, error) {
	res := make([]logistic.Product, len(products))

	for _, product := range products {
		res = append(res, product)
	}

	return res, nil
}
