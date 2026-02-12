package product

import "go-server/domain"

type ProductRepo interface {
	Create(p *domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	Find(productId int) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Delete(productId int) error
	Update(p *domain.Product) (*domain.Product, error)
	Count() (int, error)
}