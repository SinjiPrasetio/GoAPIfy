package main

import (
	"GoAPIfy/config"
	"GoAPIfy/core/helper"
	"GoAPIfy/core/service"
	"GoAPIfy/model"
	"GoAPIfy/route"
	"GoAPIfy/seeder"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	productionStr := os.Getenv("APP_PRODUCTION")
	production, err := strconv.ParseBool(productionStr)
	if err != nil {
		log.Fatal(helper.ColorizeCmd(helper.Green, "Error converting APP_PRODUCTION to boolean."))
	}

	// Print a message to indicate that the web server is being deployed
	fmt.Println(helper.ColorizeCmd(helper.Green, "Deploying Web Server..."))

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
	fmt.Println(helper.ColorizeCmd(helper.Blue, "Configuration loaded."))

	// Print a message to indicate that the server is connecting to the database
	fmt.Println(helper.ColorizeCmd(helper.Green, "Connect to database..."))

	// Create a Data Source Name (DSN) for the database connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&net_write_timeout=6000"

	// Open a connection to the database with the specified DSN and configuration
	var loggerOption = &gorm.Config{}
	// If the APP_PRODUCTION environment variable is not set to true,
	// open a connection to the database with GORM, using the provided DSN
	// and logger settings
	if !production {
		loggerOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), loggerOption)

	// Print a message to indicate that the models are being migrated
	fmt.Println(helper.ColorizeCmd(helper.Green, "Migrating models..."))

	// Automatically migrate the models to the database schema
	model.AutoMigration(db)

	// Loading modelService
	modelService := model.NewModel(db)

	seeder.RegisterSeeders(modelService)

	// Define the API routes
	fmt.Println(helper.ColorizeCmd(helper.Green, "Defining routes..."))
	route.API(server, modelService)

	// Serve static files from the public directory
	server.Static("/storage", "./public")

	// Register a handler function for the "/websocket/service" endpoint of the server
	fmt.Println("Initiating Websocket Route")
	server.GET("/websocket/service", service.WSHandler)
	fmt.Println("Websocket Deployed!")

	// Print a message to indicate that the web server is starting
	fmt.Println(helper.ColorizeCmd(helper.Green, "Starting Web Server..."))

	// Start the web server on port 8000
	server.Run(":8000")
}
