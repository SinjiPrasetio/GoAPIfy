package main

import (
	"GoAPIfy/config"
	"GoAPIfy/core/database"
	"GoAPIfy/core/helper"
	"GoAPIfy/core/service"
	"GoAPIfy/cron"
	"GoAPIfy/model"
	"GoAPIfy/route"
	"GoAPIfy/seeder"
	"GoAPIfy/service/appService"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

func main() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the configuration from the environment variables
	databaseType := os.Getenv("DATABASE_TYPE")
	productionStr := os.Getenv("APP_PRODUCTION")
	production, err := strconv.ParseBool(productionStr)
	if err != nil {
		log.Fatal(helper.ColorizeCmd(helper.Green, "Error converting APP_PRODUCTION to boolean."))
	}

	// Connecting to Meilisearch Server and initialize
	fmt.Println(helper.ColorizeCmd(helper.Green, "Connecting to Meilisearch..."))
	meilisearchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: os.Getenv("MEILI_MASTER_KEY"),
	})
	fmt.Println(helper.ColorizeCmd(helper.Green, "Meilisearch initialized..."))

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

	var db *gorm.DB
	switch databaseType {
	case "mysql", "mariadb":
		db, err = database.InitMysql(production)
	case "postgres", "postgresql":
		db, err = database.InitPostgres(production)
	case "sqlite":
		db, err = database.InitSQLite(production)
	case "mssql", "sqlserver":
		db, err = database.InitMSSQL(production)
	default:
		log.Fatal("DATABASE_TYPE is not supported. Make sure you have configure your database correctly.")
	}

	// Initialize Redis client
	fmt.Println(helper.ColorizeCmd(helper.Green, "Connecting to Redis (if enabled)..."))
	redisClient := database.InitRedis()

	// Print a message to indicate that the models are being migrated
	fmt.Println(helper.ColorizeCmd(helper.Green, "Migrating models..."))

	// Automatically migrate the models to the database schema
	model.AutoMigration(db)

	// Loading modelService
	modelService := model.NewModel(db)

	appService := appService.AppService{Model: modelService, MeiliSearch: meilisearchClient, Redis: redisClient}

	//
	fmt.Println(helper.ColorizeCmd(helper.Magenta, "Initialize Cron Jobs"))
	cron := cron.NewCron(appService)
	cron.Start()

	seeder.RegisterSeeders(appService)

	// Define the API routes
	fmt.Println(helper.ColorizeCmd(helper.Green, "Defining routes..."))
	route.API(server, appService)

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
