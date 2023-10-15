package interfaces

import "github.com/maskedemann/bookshop/internal/models"

type RepoInterface interface {
	Create(models.Book) (models.Book, error)
	Read(id string) (models.Book, error)
	Update(models.Book) error
	Delete(id string) error
}

type UsecaseInterface interface {
	Create(models.Book) (models.Book, error)
	Read(id string) (models.Book, error)
	Update(models.Book) error
	Delete(id string) error
}
