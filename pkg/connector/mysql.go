package connector

import (
	"fmt"
	"os"

	"github.com/bootmind/figo/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to the database
func Connect() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, "3306", dbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_MIGRATE") == "true" {
		Migrate(db)
	}

	return db
}

// Migrate tables
func Migrate(db *gorm.DB) {
	db.Migrator().CreateTable(&entity.Expense{})
}
