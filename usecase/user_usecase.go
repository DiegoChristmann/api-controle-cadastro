package usecase

import (
	"api-controle-cadastro/model"
	"api-controle-cadastro/repository"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (pu *UserUsecase) GetUsers() ([]model.User, error) {
	return pu.repository.GetUsers()
}

func (pu *UserUsecase) CreateUser(user model.User) (model.User, error) {

	userId, err := pu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userId

	return user, nil
}

func (pu *UserUsecase) GetUserById(id_user int) (*model.User, error) {

	user, err := pu.repository.GetUserById(id_user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (pu *UserUsecase) DeleteUser(id_user int) error {
	return pu.repository.DeleteUser(id_user)
}

func (pu *UserUsecase) UpdateUser(id_user int) error {
	return pu.repository.UpdateUser(id_user)
}
