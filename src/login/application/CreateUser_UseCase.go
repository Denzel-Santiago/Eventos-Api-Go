// Eventos-Api-Go/src/login/application/CreateUser_UseCase.go
package application

import (
	"Eventos-Api/src/login/domain"
	"Eventos-Api/src/login/domain/entities"
)

type CreateUserUseCase struct {
	userRepo domain.IUser
}

func NewCreateUserUseCase(repo domain.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo}
}

func (uc *CreateUserUseCase) Execute(user entities.User) (entities.User, error) {
	return uc.userRepo.Create(user)
}