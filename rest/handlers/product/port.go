package product

import "go-server/domain"

type Service interface {
	Find(productId int) (*domain.Product, error)
	Create(product *domain.Product) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Delete(productId int) error
	Update(product *domain.Product) (*domain.Product, error)
	Count() (int, error)
}
