package database

import (
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitSQLite initializes a connection to a SQLite database and returns a GORM instance
// If the production flag is set to false, logger settings will be applied
func InitSQLite(production bool) (*gorm.DB, error) {
	// Get the database configuration from the environment variables
	dbPath := os.Getenv("DB_PATH")

	// Open a connection to the database
	var loggerOption *gorm.Config
	if !production {
		loggerOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}, loggerOption)
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
