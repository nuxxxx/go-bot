package product

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
)

var products = map[uint64]logistic.Product{}

type ProductService interface {
	Describe(productID uint64) (*logistic.Product, error)
	List(cursor uint64, limit uint64) ([]logistic.Product, error)
	Create(logistic.Product) (uint64, error)
	Remove(productID uint64) (bool, error)
	Update(productID uint64, Product logistic.Product) error
}

type DummyProductService struct{}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}

func init() {
	var i uint64

	for i = 0; i < 100; i++ {
		products[i] = logistic.Product{
			ID:       i,
			Title:    "test",
			Price:    100,
			Quantity: 1,
		}
	}
}
