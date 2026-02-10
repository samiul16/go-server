// user/port.go
package user

import "go-server/domain"

type Service interface {
	Create(user *domain.User) (*domain.User, error)
	Find(email string) (*domain.User, error)
}

type UserRepo interface {
	Create(user *domain.User) (*domain.User, error)
	Find(email string) (*domain.User, error)
}
