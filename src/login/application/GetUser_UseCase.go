// Eventos-Api-Go/src/login/application/GetUser_UseCase.go
package application

import (
	"Eventos-Api/src/login/domain"
	"Eventos-Api/src/login/domain/entities"
)

type GetUserUseCase struct {
	userRepo domain.IUser
}

func NewGetUserUseCase(repo domain.IUser) *GetUserUseCase {
	return &GetUserUseCase{userRepo: repo}
}

func (uc *GetUserUseCase) Execute(id int) (entities.User, error) {
	return uc.userRepo.FindByID(id)
}