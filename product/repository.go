package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
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

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				 standard_cost,list_price,
				 category
				 FROM products
				 LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost,
			&product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = "SELECT COUNT(*) FROM products"
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}
