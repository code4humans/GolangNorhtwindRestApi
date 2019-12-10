package order

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getOrderByIdRequest struct {
	orderId int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

type addOrderRequest struct {
	ID           int64
	OrderDate    string
	CustomerID   int
	OrderDetails []addOrderDetailRequest
}

type addOrderDetailRequest struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
}

type deleteOrderDetailRequest struct {
	OrderDetailID string
}

type deleteOrderRequest struct {
	OrderID string
}

// @Summary Order by Id
// @Tags Order
// @Accept json
// @Produce  json
// @Param id path int true "Order Id"
// @Success 200 {object} order.OrderItem "ok"
// @Router /orders/{id} [get]
func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)
		result, err := s.GetOrderById(&req)
		helper.Catch(err)
		return result, nil

	}
	return getOrderByIdEndpoint
}

// @Summary Lista de Ordenes
// @Tags Order
// @Accept json
// @Produce  json
// @Param request body order.getOrdersRequest true "User Data"
// @Success 200 {object} order.OrderList "ok"
// @Router /orders/paginated [post]
func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	getOrdersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)

		return result, nil
	}
	return getOrdersEndpoint
}

// @Summary Insertar Order
// @Tags Order
// @Accept json
// @Produce  json
// @Param request body order.addOrderRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /orders/ [post]
func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	addOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.InsertOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return addOrderEndpoint
}

// @Summary Update Order
// @Tags Order
// @Accept json
// @Produce  json
// @Param request body order.addOrderRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /orders/ [put]
func makeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	updateOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		r, err := s.UpdateOrder(&req)
		helper.Catch(err)
		return r, nil
	}

	return updateOrderEndpoint
}

// @Summary Eliminar elemento del detalle de la Orden
// @Tags Order
// @Accept json
// @Produce  json
// @Param orderDetailId path int true "Order Detail Id"
// @Success 200 {integer} int "ok"
// @Router /orders/{orderId}/detail/{orderDetailId} [delete]
func makeDeleteOrderDetailEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDeleteEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		result, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)

		return result, nil
	}
	return deleteOrderDeleteEndpoint
}

// @Summary Eliminar Orden
// @Tags Order
// @Accept json
// @Produce  json
// @Param id path int true "Order Id"
// @Success 200 {integer} int "ok"
// @Router /orders/{id} [delete]
func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDeleteEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		r, err := s.DeleteOrder(&req)
		helper.Catch(err)
		return r, nil
	}

	return deleteOrderDeleteEndpoint
}
