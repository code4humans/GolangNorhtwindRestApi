package customer

import (
	"github.com/GolangNorhtwindRestApi/helper"
)

type Service interface {
	GetCustomers(param *getCustomersRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCustomers(param *getCustomersRequest) (*CustomerList, error) {
	customers, err := s.repo.GetCustomers(param)
	helper.Catch(err)

	totalCustomers, err := s.repo.GetTotalCustomers()
	helper.Catch(err)

	return &CustomerList{Data: customers, TotalRecords: totalCustomers}, nil
}
