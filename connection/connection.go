package connection

import (
	"accounting_transaction_api/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Connect to the database.
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka", dbConfig.DBIp, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName, dbConfig.DBPort)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

// migrate the tables.
func migrate() {
}

// GetDB returns the db instance.
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
