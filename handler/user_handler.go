package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesunonthesky/GoBackendPackage/model"
	"github.com/thesunonthesky/GoBackendPackage/repository"
)

type UserHandler struct {
	repo repository.UserRepository
}


func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Created!"})

}
