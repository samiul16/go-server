package user

import "go-server/domain"

type Service interface {
	Find(email string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}
