package controller

import (
	"api-controle-cadastro/model"
	"api-controle-cadastro/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: usecase,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	users, err := u.userUsecase.GetUsers()
	if err != nil {
		response := model.Response{
			Message: "Erro ao buscar usuários: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Garantir que sempre retorne um array vazio se não houver usuários
	// Isso evita retornar null
	if users == nil {
		users = []model.User{}
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(ctx *gin.Context) {

	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedUser, err := u.userUsecase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}

	ctx.JSON(http.StatusCreated, insertedUser)
}

func (u *UserController) GetUserById(ctx *gin.Context) {

	id := ctx.Param("userId")
	if id == "" {
		response := model.Response{
			Message: "ID do usuário nao pode ser nula",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do usuário precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.userUsecase.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := model.Response{
			Message: "O usuário nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("userId")
	if id == "" {
		response := model.Response{
			Message: "ID do usuário não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do usuário precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = u.userUsecase.DeleteUser(userId)
	if err != nil {
		// Verifica se é erro de usuário não encontrado
		if strings.Contains(err.Error(), "não encontrado") {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := model.Response{
			Message: "Erro ao deletar usuário: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Message: "Usuário deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}
