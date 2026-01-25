package main

import (
	"log"
	"myswiggyFoodDeliveryApplicaion/internal/health"
	"myswiggyFoodDeliveryApplicaion/pkg/config"
	"myswiggyFoodDeliveryApplicaion/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := database.ConnectDB(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	defer db.Close()

	r := gin.Default()

	r.GET("/health", health.Check)

	log.Println("Server running on port", cfg.AppPort)
	r.Run(":" + cfg.AppPort)

}
