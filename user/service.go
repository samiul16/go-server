package user

import "go-server/domain"

type service struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(user *domain.User) (*domain.User, error) {
	usr, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *service) Find(email string) (*domain.User, error) {
	usr, err := s.userRepo.Find(email)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

