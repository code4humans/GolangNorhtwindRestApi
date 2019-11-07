package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return productId, nil
	}

	return addProductEndpoint
}
