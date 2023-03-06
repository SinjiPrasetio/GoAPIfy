package main

import (
	"GoAPIfy/config"
	"GoAPIfy/route"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
