package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}
