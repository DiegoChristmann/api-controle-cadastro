package main

import (
	"api-controle-cadastro/controller"
	"api-controle-cadastro/repository"
	"api-controle-cadastro/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Camada de repository (usando dados mock)
	userRepository := repository.NewUserRepository(nil)
	// Camada de usecase
	usersUsecase := usecase.NewUserUsecase(userRepository)
	// Camada de controllers
	usersController := controller.NewUserController(usersUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", usersController.GetUsers)

	server.POST("/newuser", usersController.CreateUser)

	server.GET("/user/:userId", usersController.GetUserById)

	server.DELETE("/user/:userId", usersController.DeleteUser)

	server.Run(":8000")
}
