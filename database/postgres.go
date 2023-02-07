package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// initialize database connection
	postgresConnectionString := "postgresql://postgres:XzsyqhneCMlPQPrVYL3D@containers-us-west-118.railway.app:6742/railway"

	connection, err := gorm.Open(postgres.Open(postgresConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB = connection

	log.Println("Connected to database successfully!")

	// migrate models
	// err = connection.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Println(err)
	// }
	log.Println("Tables have been migrated successfully!!")
}
