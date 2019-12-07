package order

import (
	"github.com/GolangNorhtwindRestApi/helper"
)

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(param *getOrdersRequest) (*OrderList, error)
	InsertOrder(param *addOrderRequest) (int64, error)
	UpdateOrder(param *addOrderRequest) (int64, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	DeleteOrder(param *deleteOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOrderById(params *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(params)
}

func (s *service) GetOrders(param *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(param)
	helper.Catch(err)
	totalOrders, err := s.repo.GetTotalOrders(param)
	helper.Catch(err)

	return &OrderList{Data: orders, TotalRecords: totalOrders}, nil
}

func (s *service) InsertOrder(param *addOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertOrder(param)
	helper.Catch(err)

	for _, detail := range param.OrderDetails {
		detail.OrderID = orderId
		_, err := s.repo.InsertOrderDetail(&detail)
		helper.Catch(err)
	}
	return orderId, nil
}

func (s *service) UpdateOrder(param *addOrderRequest) (int64, error) {
	orderId, err := s.repo.UpdateOrder(param)
	helper.Catch(err)

	for _, detail := range param.OrderDetails {
		detail.OrderID = orderId
		if detail.ID == 0 {
			_, err = s.repo.InsertOrderDetail(&detail)
		} else {
			_, err = s.repo.UpdateOrderDetail(&detail)
		}
		helper.Catch(err)
	}
	return orderId, nil
}

func (s *service) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	return s.repo.DeleteOrderDetail(param)
}

func (s *service) DeleteOrder(param *deleteOrderRequest) (int64, error) {
	_, err := s.repo.DeleteOrderDetailByOrderId(param)
	helper.Catch(err)

	return s.repo.DeleteOrder(param)
}
