package employee

import (
	"github.com/GolangNorhtwindRestApi/helper"
)

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(param *deleteEmployeeRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)
	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)

	return &EmployeeList{
		Data:         employees,
		TotalRecords: totalEmployees,
	}, nil
}

func (s *service) GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(param)
}

func (s *service) GetBestEmployee() (*BestEmployee, error) {
	return s.repo.GetBestEmployee()
}
func (s *service) InsertEmployee(params *addEmployeeRequest) (int64, error) {
	return s.repo.InsertEmployee(params)
}

func (s *service) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	return s.repo.UpdateEmployee(params)
}

func (s *service) DeleteEmployee(param *deleteEmployeeRequest) (int64, error) {
	return s.repo.DeleteEmployee(param)
}
