// Package database provides functions for initializing and interacting with different database management systems.
package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMSSQL initializes a connection to a Microsoft SQL Server database.
// It reads the database connection details from environment variables and returns a *gorm.DB instance on success.
// If production is false, logging is enabled.
func InitMSSQL(production bool) (*gorm.DB, error) {
	// Get the database configuration from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Create the database URL string
	dbUrl := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the database
	var loggerOption *gorm.Config
	if !production {
		loggerOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(sqlserver.Open(dbUrl), &gorm.Config{}, loggerOption)
	if err != nil {
		return nil, err
	}

	// Set up database connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxIdleStr := os.Getenv("DATABASE_MAX_IDLE")
	maxIdle, err := strconv.Atoi(maxIdleStr)
	if err != nil {
		log.Fatal("Error converting DATABASE_MAX_IDLE to integer.")
	}

	maxConnectionStr := os.Getenv("DATABASE_MAX_CONNECTION")
	maxConnection, err := strconv.Atoi(maxConnectionStr)
	if err != nil {
		log.Fatal("Error converting DATABASE_MAX_CONNECTION to integer.")
	}

	maxLifetimeStr := os.Getenv("DATABASE_MAX_LIFETIME")
	maxLifetime, err := strconv.Atoi(maxLifetimeStr)
	if err != nil {
		log.Fatal("Error converting DATABASE_MAX_LIFETIME to integer.")
	}

	if maxIdle > 0 {
		sqlDB.SetMaxIdleConns(maxIdle)
	}
	if maxConnection > 0 {
		sqlDB.SetMaxOpenConns(maxConnection)
	}
	if maxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	}

	// Return the database instance
	return db, nil
}
