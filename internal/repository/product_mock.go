package repository

import "app/internal"

func NewProductsMock() *ProductsMock {
	return &ProductsMock{
		Spy: Spy{
			MethodCalls: make(map[string]int),
			MethodArgs:  make(map[string][]interface{}),
		},
	}

}

type Spy struct {
	//MethodCalls is a map of the number of times each method was called.
	MethodCalls map[string]int
	// MethodArgs is a map of the arguments that were passed to each method.
	MethodArgs map[string][]interface{}
}

type ProductsMock struct {
	//SearchProductsFunc is the function that will be called when SearchProducts is called.
	SearchProductsFunc func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	// Spy is the spy that will be used to track the method calls.
	Spy
}

// SearchProducts returns a list of products that match the query.
func (r *ProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	// increment the number of times the method was called
	r.Spy.MethodCalls["SearchProducts"]++

	// append the arguments that were passed to the method
	r.Spy.MethodArgs["SearchProducts"] = append(r.Spy.MethodArgs["SearchProducts"], query)

	// call the function that was passed to the mock
	return r.SearchProductsFunc(query)
}
