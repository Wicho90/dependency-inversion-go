package domain

type UserRepository interface {
	GetById(string) (User, error)
}
