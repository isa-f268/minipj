package main

import (
	"main/config"
	"main/model"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.DBInit()

	db.AutoMigrate(
		&model.Users{},
		&model.Topup{},
		&model.Books{},
		&model.Orders{},
		&model.Payments{},
	)
}
