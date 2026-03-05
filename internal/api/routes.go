package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func RegisterRoutes(server *gin.Engine, h *UserHandler) {
	server.LoadHTMLGlob("./templates/*")
	server.Use(favicon.New("./favicon.ico"))
	server.GET("/", h.Home)
	server.GET("/signup", h.Signup)
	server.POST("/signup", h.Signup)
	//server.GET("/:name", profile)
	//server.POST("/login", login)
}
