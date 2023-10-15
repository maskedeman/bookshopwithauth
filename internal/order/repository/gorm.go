package repository

import (
	"github.com/maskedemann/bookshop/internal/models"
	"github.com/maskedemann/bookshop/internal/order/interfaces"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(order models.Order) (models.Order, error) {
	err := repo.db.Create(&order).Error

	return order, err
}

func (repo *Repo) List(id string) (models.Order, error) {
	var order models.Order
	err := repo.db.Where("id=?", order.ID).First(&order).Error
	return order, err
}

func (repo *Repo) Delete(id string) error {
	err := repo.db.Where("id=?", id).Delete(&models.Order{}).Error
	return err
}
