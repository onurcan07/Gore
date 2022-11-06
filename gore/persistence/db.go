package persistence

import (
	"log"

	"github.com/gore/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := viper.GetString("POSTGRES_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	return db
}
