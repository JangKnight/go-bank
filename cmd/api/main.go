package main

import (
	"example.com/go-bank/internal/api"
	"example.com/go-bank/internal/db"
	"example.com/go-bank/internal/repository/postgres"
	"example.com/go-bank/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	pool, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer pool.Close()


	// User Manager
    userRepo := postgres.NewUserRepository(pool)

    // User Service
    userService := service.NewUserService(userRepo)

    // User Handler
    userHandler := api.NewUserHandler(userService)
	
	// Register Routes
	api.RegisterRoutes(server, userHandler)
	
    //Start Server...
	server.Run("0.0.0.0:8080")
}
