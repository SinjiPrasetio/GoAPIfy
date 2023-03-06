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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Deploying Web Server...")
	fmt.Println("Defining Web Server configuration...")
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOriginConfig(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	fmt.Println("Configuration loaded.")

	fmt.Println("Defining routes.")
	route.API(server)

	server.Static("/storage", "./public")
	fmt.Println("Starting Web Server...")
	server.Run(":8000")
}
