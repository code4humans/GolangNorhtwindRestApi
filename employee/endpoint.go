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

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getBestEmployeeRequest struct{}

type addEmployeeRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type updateEmployeeRequest struct {
	ID            int64
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type deleteEmployeeRequest struct {
	EmployeeID string
}

// @Summary Lista de Empleados
// @Tags Employee
// @Accept json
// @Produce  json
// @Param request body employee.getEmployeesRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "ok"
// @Router /employees/paginated [post]
func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeesEndpoint
}

// @Summary Empleado by Id
// @Tags Employee
// @Accept json
// @Produce  json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "ok"
// @Router /employees/{id} [get]
func makeGetEmployeeByIDEndpoint(s Service) endpoint.Endpoint {
	getEmployeeByIDRequest := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeeByIDRequest
}

// @Summary Mejor Empleado
// @Tags Employee
// @Accept json
// @Produce  json
// @Success 200 {object} employee.BestEmployee "ok"
// @Router /employees/best [get]
func makeGetBestEmployeeEndpoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndpoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndpoint
}

// @Summary Insertar Empleado
// @Tags Employee
// @Accept json
// @Produce  json
// @Param request body employee.addEmployeeRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [post]
func makeInsertEmployeeEndpoint(s Service) endpoint.Endpoint {
	getInsertEmployeeEndpoint := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		result, err := s.InsertEmployee(&req)
		helper.Catch(err)
		return result, nil
	}
	return getInsertEmployeeEndpoint
}

// @Summary Update Empleado
// @Tags Employee
// @Accept json
// @Produce  json
// @Param request body employee.updateEmployeeRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [put]
func makeUpdateEmployeeEndpoint(s Service) endpoint.Endpoint {
	getUpdateEmployeeEndpoint := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		result, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return result, nil
	}
	return getUpdateEmployeeEndpoint
}

// @Summary Eliminar Empleado
// @Tags Employee
// @Accept json
// @Produce  json
// @Param id path int true "Employee Id"
// @Success 200 {integer} int "ok"
// @Router /employees/{id} [delete]
func makeDeleteEmployeeEndpoint(s Service) endpoint.Endpoint {
	getDeleteEmployeeEndpoint := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		result, err := s.DeleteEmployee(&req)
		helper.Catch(err)

		return result, nil
	}
	return getDeleteEmployeeEndpoint
}
