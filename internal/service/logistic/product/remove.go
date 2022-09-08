package product

func (ps *DummyProductService) Remove(productID uint64) (bool, error) {
	delete(products, productID)

	return true, nil
}
