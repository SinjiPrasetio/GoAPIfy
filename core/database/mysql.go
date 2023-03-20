package database

import (
	"GoAPIfy/core/helper"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMysql initializes a new MySQL or MariaDB database connection and returns a GORM DB object, or an error if the connection fails.
// If the production flag is set to true, GORM will not output SQL logs.
func InitMysql(production bool) (*gorm.DB, error) {
	// Print a message to indicate that the MySQL / MariaDB connection is being initialized
	fmt.Println(helper.ColorizeCmd(helper.Blue, "Initializing MySQL / MariaDB"))

	// Get the database configuration from the environment variables
	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")

	// Create a Data Source Name (DSN) for the database connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&net_write_timeout=6000"

	// Set the logger configuration based on the production flag
	var loggerOption = &gorm.Config{}
	if !production {
		loggerOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	// Open a connection to the database using the provided DSN and logger settings
	db, err := gorm.Open(mysql.Open(dsn), loggerOption)
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
	return db, err
}
