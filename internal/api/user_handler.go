package api

import (
	"log"
	"net/http"

	"example.com/go-bank/internal/service"
	"github.com/gin-gonic/gin"
	// 	"strings"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Home(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}

func (h *UserHandler) Signup(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "signup.html", nil)
		return
	}

	var input struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=4"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.svc.RegisterUser(c.Request.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		log.Printf("Service Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

