package repository

import "app/internal"

func NewProductsMapStub() *ProductsMapStub {
	return &ProductsMapStub{}
}

type ProductsMapStub struct {
	//SearchProductsFunc is the function that will be called when SearchProducts is called.
	SearchProductsFunc func(query internal.ProductQuery) (p map[int]internal.Product, err error)
}

// SearchProducts returns a list of products that match the query.
func (r *ProductsMapStub) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	return r.SearchProductsFunc(query)
}
