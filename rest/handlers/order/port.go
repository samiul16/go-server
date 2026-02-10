package order

import "go-server/domain"

type Service interface {
	Create(order *domain.Order) (*domain.Order, error)
}