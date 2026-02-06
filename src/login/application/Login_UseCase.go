// Eventos-Api-Go/src/login/application/Login_UseCase.go
package application

import (
	"Eventos-Api/src/login/domain"
	"Eventos-Api/src/login/domain/entities"
)

type LoginUseCase struct {
	userRepo domain.IUser
}

func NewLoginUseCase(repo domain.IUser) *LoginUseCase {
	return &LoginUseCase{userRepo: repo}
}

func (uc *LoginUseCase) Execute(username, password string) (entities.User, error) {
	return uc.userRepo.Authenticate(username, password)
}