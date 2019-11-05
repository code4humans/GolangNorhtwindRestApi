package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				standard_cost,list_price,
				category
				FROM products
				WHERE id=?`
	row := repo.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description,
		&product.StandardCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}
	return product, err
}
