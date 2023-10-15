package migration

import (
	"fmt"

	db "github.com/maskedemann/bookshop/internal/conf"
	"github.com/maskedemann/bookshop/internal/models"
	"gorm.io/gorm"
)

func main() {
	conn := db.DBConn()
	MakeMigrations(conn)
	return
}

func MakeMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("Migration failed for User model.")
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		fmt.Println("Migration failed for Book model.")
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		fmt.Println("Migration failed for Order model.")
		panic(err.Error())
	}
}
