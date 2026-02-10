package product

import "go-server/domain"

type Service interface {
	Create(product *domain.Product) (*domain.Product, error)
	Find(id int) (*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(productId int) error
	List() ([]*domain.Product, error)
}

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (s *service) Create(product *domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(product)
}

func (s *service) Find(id int) (*domain.Product, error) {
	return s.productRepo.Find(id)		
}

func (s *service) Update(product *domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(product)
}

func (s *service) Delete(productId int) error {
	return s.productRepo.Delete(productId)
}


func (s *service) List() ([]*domain.Product, error) {
	return s.productRepo.List()
}

