package main

import (
	"fmt"
	"log"
	"meeting-scheduler/config"
	"meeting-scheduler/handlers"
	"meeting-scheduler/repositories"

	"github.com/gin-gonic/gin"
)

func main() {

	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := repositories.InitDB(configuration.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	handlers.RegisterRoutes(router, db)
	config.CreateTable(configuration.SqlFilePath,db)
	port := configuration.ServerPort
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(router.Run(":" + port))

}
