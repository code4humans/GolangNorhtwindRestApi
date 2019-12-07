package order

type OrderList struct {
	Data         []*OrderItem `json:"data"`
	TotalRecords int64        `json:"totalRecords"`
}

type OrderItem struct {
	ID         int64              `json:"orderId"`
	CustomerID int                `json:"customerId"`
	OrderDate  string             `json:"orderDate"`
	StatusId   string             `json:"statusId"`
	StatusName string             `json:"statusName"`
	Customer   string             `json:"customer"`
	Company    string             `json:"company"`
	Address    string             `json:"address"`
	Phone      string             `json:"phone"`
	City       string             `json:"city"`
	Data       []*OrderDetailItem `json:"data"`
}

type OrderDetailItem struct {
	ID          int64   `json:"id"`
	OrderId     int     `json:"order_id"`
	ProductId   int     `json:"product_id"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	ProductName string  `json:"product_name"`
}
