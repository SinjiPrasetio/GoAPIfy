package main

import (
	"GoAPI/model"
	"GoAPIfy/config"
	"GoAPIfy/route"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database configuration from the environment variables
	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")

	// Print a message to indicate that the web server is being deployed
	fmt.Println("\033[34mDeploying Web Server...\033[0m")

	// Create a new gin server with default middleware
	server := gin.Default()

	// Set up Cross-Origin Resource Sharing (CORS)
	server.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOriginConfig(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Print a message to indicate that the server configuration has been loaded
	fmt.Println("\033[32mConfiguration loaded.\033[0m")

	// Print a message to indicate that the server is connecting to the database
	fmt.Println("\033[34mConnect to database...\033[0m")

	// Create a Data Source Name (DSN) for the database connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&net_write_timeout=6000"

	// Open a connection to the database with the specified DSN and configuration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Uncomment below to enable logging of database transactions.
		// Logger: logger.Default.LogMode(logger.Info),
	})

	// Print a message to indicate that the models are being migrated
	fmt.Println("\033[34mMigrating models...\033[0m")

	// Automatically migrate the models to the database schema
	model.AutoMigration(db)

	// Define the API routes
	fmt.Println("\033[34mDefining routes...\033[0m")
	route.API(server)

	// Serve static files from the public directory
	server.Static("/storage", "./public")

	// Print a message to indicate that the web server is starting
	fmt.Println("\033[34mStarting Web Server...\033[0m")

	// Start the web server on port 8000
	server.Run(":8000")
}
