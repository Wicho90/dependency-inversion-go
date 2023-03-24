package infrastructure

import (
	"errors"

	"github.com/wicho90/dependency-inversion-go/domain"
)

var users domain.Users = []domain.User{
	{Id: "1", Email: "occhoa90@gmail.com"},
	{Id: "2", Email: "alberto@gmail.com"},
	{Id: "3", Email: "victor@gmail.com"},
	{Id: "4", Email: "karla@gmail.com"},
}

type InMemoryUserRepository struct {
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}
func (i *InMemoryUserRepository) GetById(id string) (domain.User, error) {

	for _, u := range users {
		if id == u.Id {
			return u, nil
		}
	}

	return domain.User{}, errors.New("Usuario no encontrado")
}
