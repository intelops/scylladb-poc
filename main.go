package main

import (
	"log"

	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/controllers"
	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/daos/clients/scylla"
	"github.com/gin-gonic/gin"
	"github.com/scylladb/gocqlx/v2"
)

var (
	scyllaManager *scylla.Manager
	scyllaErr     error
	scyllaSession gocqlx.Session
)

func init() {
	scyllaManager, scyllaErr = scylla.NewManager()
	if scyllaErr != nil {
		log.Fatalf("error occurred while reading scylla configs: %s", scyllaErr)
		return
	}
	scyllaSession, scyllaErr = scyllaManager.Connect()
	if scyllaErr != nil {
		log.Fatalf("error occurred while connecting to scylla: %s", scyllaErr)
		return
	}
	log.Println("Connected to scylla successfully...")
	defer scyllaSession.Close()
}

func main() {
	server := gin.Default()
	// router := server.Group("/api/v1")
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Route not found"})
	})
	UserController, err := controllers.NewUserController()
	if err != nil {
		log.Fatalf("error occurred while creating user controller: %s", err)
		return
	}
	v1 := server.Group("/api/v1")
	{
		v1.POST("/user", UserController.CreateUser)
		v1.GET("/user/:id", UserController.GetUser)
		v1.GET("/users", UserController.GetUsers)
		v1.PUT("/user/:id", UserController.UpdateUser)
		v1.DELETE("/user/:id", UserController.DeleteUser)
	}
	if err = server.Run(":8080"); err != nil {
		log.Fatalf("error occurred while running server: %s", err)
		return
	}
}
