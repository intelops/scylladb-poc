package controllers

import (
	"log"
	"net/http"

	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/models"
	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/services"
	"github.com/MrAzharuddin/scylladb-gin/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() (*UserController, error) {
	userService, err := services.NewUserService()
	if err != nil {
		return nil, err
	}
	return &UserController{UserService: userService}, nil
}

func (uc *UserController) CreateUser(context *gin.Context) {
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	input.Id = utils.GenerateUUID().String()
	if _, err := uc.UserService.CreateUser(&input); err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (uc *UserController) GetUser(context *gin.Context) {
	id := context.Param("id")
	user, err := uc.UserService.GetUser(id)
	if err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uc *UserController) GetUsers(context *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uc *UserController) UpdateUser(context *gin.Context) {
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := uc.UserService.UpdateUser(&input); err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (uc *UserController) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	if err := uc.UserService.DeleteUser(id); err != nil {
		log.Println(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}