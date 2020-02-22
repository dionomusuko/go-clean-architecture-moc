package usecase

import "github.com/dionomusuko/go-crean-archi/domain"

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindById(id int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(domain.User) (domain.User, error)
	Delete(domain.User) (domain.User, error)
}
