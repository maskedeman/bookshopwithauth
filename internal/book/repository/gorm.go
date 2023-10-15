package repository

import (
	"github.com/maskedemann/bookshop/internal/app/Book/interfaces"
	"github.com/maskedemann/bookshop/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(book models.Book) (models.Book, error) {
	err := repo.db.Create(&book).Error
	return book, err
}

func (repo *Repo) Read(id string) (models.Book, error) {
	var book models.Book
	err := repo.db.Where("id=?", book.ID).First(&book).Error
	return book, err
}
func (repo *Repo) Update(book models.Book) error {

	var dbBook models.Book
	if err := repo.db.Where("id=?", book.ID).First(&dbBook).Error; err != nil {
		return err
	}
	dbBook.Name = book.Name
	dbBook.Author = book.Author
	dbBook.Publication = book.Publication
	dbBook.Description = book.Description
	dbBook.Isbn = book.Isbn
	err := repo.db.Save(dbBook).Error
	return err

}
func (repo *Repo) Delete(id string) error {
	err := repo.db.Where("id=?", id).Delete(&models.Book{}).Error
	return err
}
