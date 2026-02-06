// Eventos-Api-Go/src/login/application/DeleteUser_UseCase.go
package application

import (
	"Eventos-Api/src/login/domain"
)

type DeleteUserUseCase struct {
	userRepo domain.IUser
}

func NewDeleteUserUseCase(repo domain.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: repo}
}

func (uc *DeleteUserUseCase) Execute(id int) error {
	return uc.userRepo.Delete(id)
}