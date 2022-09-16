package product

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

func (ps *DummyProductService) List(cursor uint64, limit uint64) ([]logistic.Product, error) {
	res := make([]logistic.Product, limit+1)

	start := (cursor - 1) * limit

	for i := 0; i <= int(limit); i++ {
		res[i] = products[start+uint64(i)]
	}

	return res, nil
}
