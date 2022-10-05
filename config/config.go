package config

import (
	"fmt"
	"log"
	"praktek1/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	config := fmt.Sprintf("host=localhost user=postgres password=alam dbname=praktek1 port=5432 sslmode=disable")

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(structs.Person{})
	return db
}
