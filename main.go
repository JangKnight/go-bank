package main

import (
	"example.com/go-bank/db"
	"example.com/go-bank/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run("0.0.0.0:8080")
}
