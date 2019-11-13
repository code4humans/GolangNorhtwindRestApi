package employee

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeesEndpoint
}
