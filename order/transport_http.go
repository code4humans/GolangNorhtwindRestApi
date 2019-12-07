package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GolangNorhtwindRestApi/helper"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdHandler := kithttp.NewServer(makeGetOrderByIdEndpoint(s),
		getOrderByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	getOrdersHandler := kithttp.NewServer(makeGetOrdersEndpoint(s),
		getOrdersRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	addOrderHandler := kithttp.NewServer(makeAddOrderEndpoint(s),
		addOrderRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addOrderHandler)

	updateOrderHandler := kithttp.NewServer(makeUpdateOrderEndpoint(s),
		getUpdateOrderRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateOrderHandler)

	deleteOrderDetailHandler := kithttp.NewServer(makeDeleteOrderDetailEndpoint(s),
		getDeleteOrderDetailRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{orderId}/detail/{orderDetailId}", deleteOrderDetailHandler)

	deleteOrderHandler := kithttp.NewServer(makeDeleteOrderEndpoint(s),
		getDeleteOrderRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteOrderHandler)

	return r
}

func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)
	return getOrderByIdRequest{
		orderId: orderId,
	}, nil
}

func getOrdersRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func addOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func getUpdateOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getDeleteOrderDetailRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderDetailRequest{
		OrderDetailID: chi.URLParam(r, "orderDetailId"),
	}, nil
}

func getDeleteOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderRequest{
		OrderID: chi.URLParam(r, "id"),
	}, nil
}
