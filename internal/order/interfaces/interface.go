package interfaces

import "github.com/maskedemann/bookshop/internal/models"

type RepoInterface interface {
	Create(models.Order) (models.Order, error)
	List(id string) (models.Order, error)
	Delete(id string) error
}

type UsecaseInterface interface {
	Create(models.Order) (models.Order, error)
	List(id string) (models.Order, error)
	Delete(id string) error
}
