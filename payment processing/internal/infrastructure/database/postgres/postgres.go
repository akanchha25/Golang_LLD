package postgres

import (
	"payment_processing/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=password "), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Payment{})
	DB = db
}
