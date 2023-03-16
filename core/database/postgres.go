// Package database provides functions for initializing and connecting to a PostgreSQL database using the GORM ORM library.
package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitPostgres initializes a connection to a PostgreSQL database and returns a GORM DB instance. The function takes a boolean argument production which determines whether to use a logger in production environment or not.
// It reads database configuration from environment variables: DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS, DATABASE_MAX_IDLE, DATABASE_MAX_CONNECTION, and DATABASE_MAX_LIFETIME.
// It creates the database URL string using the configuration and opens a connection to the database. Then, it sets up the database connection pool settings.
// The function returns a GORM DB instance and an error if any.
func InitPostgres(production bool) (*gorm.DB, error) {

	// Get the database configuration from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Create the database URL string
	dbUrl := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)

	// Open a connection to the database
	var loggerOption *gorm.Config
	if !production {
		loggerOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(postgres.Open(dbUrl), loggerOption)
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

	// Configure database connection pool
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
