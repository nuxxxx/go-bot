package logistic

import (
	"fmt"
)

type Product struct {
	ID       uint64
	Title    string
	Price    float64
	Quantity int
}

func (p *Product) String() string {
	return fmt.Sprintf("Product: \n\nID: %v\nTitle: %v\nPrice: %v$\nQuantity: %v", p.ID, p.Title, p.Price, p.Quantity)
}
