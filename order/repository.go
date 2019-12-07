package order

import (
	"database/sql"
	"fmt"

	"github.com/GolangNorhtwindRestApi/helper"
)

type Repository interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(param *getOrdersRequest) ([]*OrderItem, error)
	GetTotalOrders(param *getOrdersRequest) (int64, error)
	InsertOrder(param *addOrderRequest) (int64, error)
	InsertOrderDetail(param *addOrderDetailRequest) (int64, error)
	UpdateOrder(param *addOrderRequest) (int64, error)
	UpdateOrderDetail(param *addOrderDetailRequest) (int64, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	DeleteOrderDetailByOrderId(param *deleteOrderRequest) (int64, error)
	DeleteOrder(param *deleteOrderRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetOrderById(param *getOrderByIdRequest) (*OrderItem, error) {
	var sql = `SELECT o.id,o.customer_id,o.order_date,o.status_id,os.status_name,
	CONCAT(c.first_name,' ',c.last_name) as customer_name,
	c.company,
	c.address,
	c.business_phone,
	c.city
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE o.id = ? `

	order := &OrderItem{}

	row := repo.db.QueryRow(sql, param.orderId)
	err := row.Scan(&order.ID, &order.CustomerID, &order.OrderDate,
		&order.StatusId, &order.StatusName, &order.Customer,
		&order.Company, &order.Address, &order.Phone, &order.City)
	helper.Catch(err)
	orderDetail, err := GetOrderDetail(repo, &param.orderId)
	helper.Catch(err)
	order.Data = orderDetail

	return order, nil
}

func GetOrderDetail(repo *repository, orderId *int64) ([]*OrderDetailItem, error) {
	const sql = `SELECT order_id,od.id,quantity,unit_price,p.product_name,product_id
	FROM order_details od
	INNER JOIN products p ON od.product_id = p.id
	WHERE od.order_id = ?`

	results, err := repo.db.Query(sql, orderId)
	helper.Catch(err)

	var orders []*OrderDetailItem
	for results.Next() {
		order := &OrderDetailItem{}
		err = results.Scan(&order.OrderId, &order.ID, &order.Quantity,
			&order.UnitPrice, &order.ProductName, &order.ProductId)
		helper.Catch(err)
		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *repository) GetOrders(param *getOrdersRequest) ([]*OrderItem, error) {
	var filter string

	if param.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id  = %v ", param.Status.(float64))
	}

	if param.DateFrom != nil && param.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date  >= '%v' ", param.DateFrom.(string))
	}

	if param.DateFrom == nil && param.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date  <= '%v' ", param.DateTo.(string))
	}

	if param.DateFrom != nil && param.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date  between '%v' and '%v' ", param.DateFrom.(string), param.DateTo.(string))
	}
	var sql = `SELECT o.id,
	o.customer_id,
	o.order_date,
	o.status_id,
	os.status_name,
	CONCAT(c.first_name,' ',c.last_name) as customer_name	
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE 1=1 `

	sql = sql + filter + "LIMIT ? OFFSET ?"

	results, err := repo.db.Query(sql, param.Limit, param.Offset)
	helper.Catch(err)

	var orders []*OrderItem

	for results.Next() {
		order := &OrderItem{}
		err = results.Scan(&order.ID, &order.CustomerID, &order.OrderDate,
			&order.StatusId, &order.StatusName, &order.Customer)
		helper.Catch(err)

		orderDetail, err := GetOrderDetail(repo, &order.ID)
		helper.Catch(err)

		order.Data = orderDetail
		orders = append(orders, order)
	}
	return orders, err
}

func (repo *repository) GetTotalOrders(param *getOrdersRequest) (int64, error) {
	var filter string

	if param.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id  = %v ", param.Status.(float64))
	}

	if param.DateFrom != nil && param.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date  >= '%v' ", param.DateFrom.(string))
	}

	if param.DateFrom == nil && param.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date  <= '%v' ", param.DateTo.(string))
	}

	if param.DateFrom != nil && param.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date  between '%v' and '%v' ", param.DateFrom.(string), param.DateTo.(string))
	}

	var sql = "SELECT COUNT(*) FROM orders o WHERE 1=1 " + filter
	row := repo.db.QueryRow(sql)

	var total int64
	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func (repo *repository) InsertOrder(param *addOrderRequest) (int64, error) {
	const sql = `
	INSERT INTO orders
	(customer_id ,order_date)
	VALUES(?,?)`

	result, err := repo.db.Exec(sql, param.CustomerID, param.OrderDate)
	helper.Catch(err)

	id, err := result.LastInsertId()
	helper.Catch(err)
	return id, nil
}

func (repo *repository) InsertOrderDetail(param *addOrderDetailRequest) (int64, error) {
	const sql = `
	INSERT INTO order_details
	(order_id ,product_id,quantity,unit_price)
	VALUES(?,?,?,?)`

	result, err := repo.db.Exec(sql, param.OrderID, param.ProductID, param.Quantity, param.UnitPrice)
	helper.Catch(err)

	id, err := result.LastInsertId()
	helper.Catch(err)
	return id, nil
}

func (repo *repository) UpdateOrder(param *addOrderRequest) (int64, error) {
	const sql = `
	UPDATE orders
	SET customer_id = ?
	WHERE id = ? `

	_, err := repo.db.Exec(sql, param.CustomerID, param.ID)
	helper.Catch(err)

	return param.ID, nil
}

func (repo *repository) UpdateOrderDetail(param *addOrderDetailRequest) (int64, error) {
	const sql = `
	UPDATE order_details
	SET quantity = ?,
		unit_price = ?
	WHERE id = ?`

	_, err := repo.db.Exec(sql, param.Quantity, param.UnitPrice, param.ID)
	helper.Catch(err)

	return param.ID, nil
}

func (repo *repository) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE id = ?`
	result, err := repo.db.Exec(sql, param.OrderDetailID)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)

	return count, nil
}

func (repo *repository) DeleteOrderDetailByOrderId(param *deleteOrderRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE order_id = ?`
	result, err := repo.db.Exec(sql, param.OrderID)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}

func (repo *repository) DeleteOrder(param *deleteOrderRequest) (int64, error) {
	const sql = `DELETE FROM orders WHERE id = ?;`
	result, err := repo.db.Exec(sql, param.OrderID)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}
