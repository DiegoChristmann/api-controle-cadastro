package main

import (
	"api-controle-cadastro/controller"
	"api-controle-cadastro/db"
	"api-controle-cadastro/repository"
	"api-controle-cadastro/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer database.Close()

	server := gin.Default()

	// Camada de repository (usando banco de dados PostgreSQL)
	userRepository := repository.NewUserRepository(database)
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
