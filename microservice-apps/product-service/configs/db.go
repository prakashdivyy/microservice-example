package configs

import (
	"fmt"
	"os"

	"microservice/product-service/products"

	"github.com/jinzhu/gorm"
)

// InitDB will initialize DB connection
func InitDB() *gorm.DB {
	dbType := "mysql"
	if os.Getenv("DB_TYPE") != "" {
		dbType = os.Getenv("DB_TYPE")
	}

	dbUsername := "root"
	if os.Getenv("DB_USERNAME") != "" {
		dbUsername = os.Getenv("DB_USERNAME")
	}

	dbPassword := "root"
	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}

	dbHost := "localhost"
	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	}

	dbPort := "3306"
	if os.Getenv("DB_PORT") != "" {
		dbPort = os.Getenv("DB_PORT")
	}

	dbName := "product_service"
	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	}

	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(dbType, dbURL)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&products.Product{})

	return db
}
