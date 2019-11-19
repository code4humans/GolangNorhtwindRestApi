package customer

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

func makeGetcustomersEndpoint(s Service) endpoint.Endpoint {
	getCustomersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)

		return result, nil
	}
	return getCustomersEndpoint
}
