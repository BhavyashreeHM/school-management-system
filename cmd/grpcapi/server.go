package main

import (
	"employee-management-system/internal/reposioriy/mongodb"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	port := os.Getenv("APP_PORT")

	_, err = mongodb.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	r := gin.Default()
	log.Println("Server running on port", port)
	r.Run(":" + port)

}
