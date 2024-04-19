package database

import (
	"fmt"
	"log"
	"os"

	schema "github.com/MalshanPerera/go-expense-tracker/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func createConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
}

func Connect() {
	connStr := createConnStr()
	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = schema.DBMigrate(gormDB)

	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	log.Println("Connected to the database")

	db = gormDB
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close the database: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close the database: %v", err)
	}

	log.Println("Database connection closed")
}

func GetDB() *gorm.DB {
	return db
}
