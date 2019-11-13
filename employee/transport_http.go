package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/GolangNorhtwindRestApi/helper"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndpoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeByIdHandler := kithttp.NewServer(makeGetEmployeeByIDEndpoint(s),
		getEmployeeByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIdHandler)

	getBestEmployeeHandler := kithttp.NewServer(makeGetBestEmployeeEndpoint(s),
		getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeHandler)
	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeeByIDRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}

func getBestEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}
