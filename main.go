package main

import (
	"net/http"

	"github.com/GolangNorhtwindRestApi/customer"
	"github.com/GolangNorhtwindRestApi/database"
	_ "github.com/GolangNorhtwindRestApi/docs"
	"github.com/GolangNorhtwindRestApi/employee"
	"github.com/GolangNorhtwindRestApi/order"
	"github.com/GolangNorhtwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Northwind API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository  = product.NewRepository(databaseConnection)
		employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
		orderRepository    = order.NewRepository(databaseConnection)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	r.Mount("/orders", order.MakeHTTPHandler(orderService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)
}
