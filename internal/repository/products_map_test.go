package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProducts(t *testing.T) {
	t.Run("case 01 - search products by id", func(t *testing.T) {
		// arrange
		// initialize the products map
		products := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		// create the products map
		productsMap := repository.NewProductsMap(products)

		// act
		products, err := productsMap.SearchProducts(internal.ProductQuery{
			Id: 1,
		})
		require.NoError(t, err)

		expectedProduct := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
		}

		// assert
		require.Equal(t, expectedProduct, products)

	})
}
