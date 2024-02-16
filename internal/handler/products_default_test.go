package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

type testCases struct {
	name                       string
	searchProducts             func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	method                     string
	url                        string
	expectedStatusCode         int
	expectedHeader             http.Header
	expectedBody               string
	expectedCallSearchProducts int
}

func TestGetWithSliceCaseTest(t *testing.T) {
	cases := []testCases{
		{
			name: "case 01 - return all products",
			searchProducts: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
				p = map[int]internal.Product{
					1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "product 1", Price: 10, SellerId: 1}},
					2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "product 2", Price: 20, SellerId: 2}},
					3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "product 3", Price: 30, SellerId: 3}},
					4: {Id: 4, ProductAttributes: internal.ProductAttributes{Description: "product 4", Price: 40, SellerId: 4}},
					5: {Id: 5, ProductAttributes: internal.ProductAttributes{Description: "product 5", Price: 50, SellerId: 5}},
				}
				return
			},
			method:                     "GET",
			url:                        "/products",
			expectedStatusCode:         http.StatusOK,
			expectedHeader:             http.Header{"Content-Type": []string{"application/json"}},
			expectedBody:               `{"message": "success","data":{"1":{"id":1, "description":"product 1", "price":10, "seller_id":1},"2":{"id":2, "description":"product 2", "price":20, "seller_id":2},"3":{"id":3, "description":"product 3", "price":30, "seller_id":3},"4":{"id":4, "description":"product 4", "price":40, "seller_id":4},"5":{"id":5, "description":"product 5", "price":50, "seller_id":5}}}`,
			expectedCallSearchProducts: 1,
		},
		{
			name: "case 02 - return product by id",
			searchProducts: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
				p = map[int]internal.Product{
					1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "product 1", Price: 10, SellerId: 1}},
				}
				return
			},
			method:                     "GET",
			url:                        "/products?id=1",
			expectedStatusCode:         http.StatusOK,
			expectedHeader:             http.Header{"Content-Type": []string{"application/json"}},
			expectedBody:               `{"message": "success","data":{"1":{"id":1, "description":"product 1", "price":10, "seller_id":1}}}`,
			expectedCallSearchProducts: 1,
		},
		{
			name:                       "case 03 - invalid id",
			searchProducts:             nil,
			method:                     "GET",
			url:                        "/products?id=1a",
			expectedStatusCode:         http.StatusBadRequest,
			expectedHeader:             http.Header{"Content-Type": []string{"application/json"}},
			expectedBody:               `{"status": "Bad Request", "message": "invalid id"}`,
			expectedCallSearchProducts: 0,
		},
		{
			name: "case 04 - internal error",
			searchProducts: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
				err = errors.New("internal error")
				return
			},
			method:                     "GET",
			url:                        "/products",
			expectedStatusCode:         http.StatusInternalServerError,
			expectedHeader:             http.Header{"Content-Type": []string{"application/json"}},
			expectedBody:               `{"message": "internal error","status": "Internal Server Error"}`,
			expectedCallSearchProducts: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// repository: mock
			rp := repository.NewProductsMock()

			// set searchProducts func
			if tc.searchProducts != nil {
				rp.SearchProductsFunc = tc.searchProducts
			}

			// handler
			hd := handler.NewProductsDefault(rp)

			// request
			req := httptest.NewRequest(tc.method, tc.url, nil)

			// response
			res := httptest.NewRecorder()

			// act
			hd.Get()(res, req)

			// assert
			require.Equal(t, tc.expectedStatusCode, res.Code)
			require.Equal(t, tc.expectedHeader, res.Header())
			require.JSONEq(t, tc.expectedBody, res.Body.String())
			require.Equal(t, tc.expectedCallSearchProducts, rp.Spy.MethodCalls["SearchProducts"])
		})

	}
}

func TestGet(t *testing.T) {
	t.Run("case 01 - return all products", func(t *testing.T) {
		// arrange
		// repository: mock
		rp := repository.NewProductsMock()
		rp.SearchProductsFunc = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = map[int]internal.Product{
				1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "product 1", Price: 10, SellerId: 1}},
				2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "product 2", Price: 20, SellerId: 2}},
				3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "product 3", Price: 30, SellerId: 3}},
				4: {Id: 4, ProductAttributes: internal.ProductAttributes{Description: "product 4", Price: 40, SellerId: 4}},
				5: {Id: 5, ProductAttributes: internal.ProductAttributes{Description: "product 5", Price: 50, SellerId: 5}},
			}

			return
		}
		// handler
		hd := handler.NewProductsDefault(rp)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()

		// act
		hd.Get()(res, req)

		// assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message": "success",
		"data":{
			"1":{"id":1, "description":"product 1", "price":10, "seller_id":1},
			"2":{"id":2, "description":"product 2", "price":20, "seller_id":2},
			"3":{"id":3, "description":"product 3", "price":30, "seller_id":3},
			"4":{"id":4, "description":"product 4", "price":40, "seller_id":4},
			"5":{"id":5, "description":"product 5", "price":50, "seller_id":5}
		}
		}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, 1, rp.Spy.MethodCalls["SearchProducts"])
	})

	t.Run("case 02 - return a product by id", func(t *testing.T) {
		// arrange
		// repository: mock
		rp := repository.NewProductsMock()
		rp.SearchProductsFunc = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = map[int]internal.Product{
				1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "product 1", Price: 10, SellerId: 1}},
			}
			return
		}
		// handler
		hd := handler.NewProductsDefault(rp)

		r := chi.NewRouter()
		r.Get("/products", hd.Get())

		req, _ := http.NewRequest("GET", "/products?id=1", nil)
		res := httptest.NewRecorder()

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message": "success",
		"data":{
			"1":{"id":1, "description":"product 1", "price":10, "seller_id":1}
		}
		}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, 1, rp.Spy.MethodCalls["SearchProducts"])
	})

	t.Run("case 03 - invalid id", func(t *testing.T) {
		// arrange
		// repository: mock
		rp := repository.NewProductsMock()
		// handler
		hd := handler.NewProductsDefault(rp)

		req := httptest.NewRequest("GET", "/products?id=a", nil)
		res := httptest.NewRecorder()

		// act
		hd.Get()(res, req)

		// assert
		expectedStatusCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"status": "Bad Request", "message": "invalid id"}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, 0, rp.Spy.MethodCalls["SearchProducts"])
	})

	t.Run("case 04 - internal server error", func(t *testing.T) {
		// arrange
		// repository: mock
		rp := repository.NewProductsMock()
		rp.SearchProductsFunc = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			err = errors.New("unkown error")
			return
		}
		// handler
		hd := handler.NewProductsDefault(rp)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()

		// act
		hd.Get()(res, req)

		// assert
		expectedStatusCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message": "internal error","status": "Internal Server Error"}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, 1, rp.Spy.MethodCalls["SearchProducts"])
	})
}
