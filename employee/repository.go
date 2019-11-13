package employee

import (
	"database/sql"

	"github.com/GolangNorhtwindRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT 
				id, first_name, last_name,
				company, email_address, job_title,
				business_phone,home_phone,
				COALESCE(mobile_phone,''),fax_number,address
 				FROM employees
				 LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err = results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
			&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
			&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch(err)
		employees = append(employees, employee)
	}

	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = "SELECT COUNT(*) FROM employees"
	var total int64

	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func (repo *repository) GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error) {
	const sql = `SELECT id,first_name,last_name,company,email_address,
	             job_title,business_phone,home_phone,
				COALESCE(mobile_phone,''),fax_number,address
				FROM employees
				WHERE id=?`
	row := repo.db.QueryRow(sql, param.EmployeeID)
	employee := &Employee{}
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
		&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
		&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
	helper.Catch(err)

	return employee, nil
}

func (repo *repository) GetBestEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id,count(e.id) as totalVentas,e.first_name,e.last_name
				FROM orders o
				INNER JOIN employees e  ON o.employee_id = e.id
				GROUP BY o.employee_id
				ORDER BY totalVentas desc
				limit 1`

	row := repo.db.QueryRow(sql)
	employee := &BestEmployee{}
	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.FirstName, &employee.LastName)
	helper.Catch(err)

	return employee, nil

}
