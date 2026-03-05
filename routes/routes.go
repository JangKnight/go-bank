package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.LoadHTMLGlob("./templates/*")
	server.GET("/", home)
	//server.GET("/:name", profile)
	//server.POST("/signup", signup)
	//server.POST("/login", login)
}
