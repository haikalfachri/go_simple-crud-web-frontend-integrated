package database

import (
	"biodata/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}
  
func (c *Config) ConnectDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  c.DB_USERNAME,
	  c.DB_PASSWORD,
	  c.DB_HOST,
	  c.DB_PORT,
	  c.DB_NAME,
	)
	
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
	  panic(err)
	}

	log.Printf("successfully connected to database\n")
}

func MigrateDB() {

	err := DB.AutoMigrate(&models.Biodata{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

	log.Printf("successfully database migration\n")
}

