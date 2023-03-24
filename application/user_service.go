package application

import (
	"github.com/wicho90/dependency-inversion-go/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) FinById(id string) (domain.User, error) {

	user, err := s.repository.GetById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
