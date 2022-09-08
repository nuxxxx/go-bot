package product

func (ps *DummyProductService) ProductsLength() uint64 {
	return uint64(len(products))
}
