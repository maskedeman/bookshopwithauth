package interfaces

import "github.com/maskedemann/bookshop/internal/models"

type RepoInterface interface {
	Create(models.User) (models.User, error)
	Read(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}

type UsecaseInterface interface {
	Create(models.User) (models.User, error)
	Read(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}
