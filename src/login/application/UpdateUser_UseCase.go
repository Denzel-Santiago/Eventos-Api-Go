// Eventos-Api-Go/src/login/application/UpdateUser_UseCase.go
package application

import (
	"Eventos-Api/src/login/domain"
	"Eventos-Api/src/login/domain/entities"
)

type UpdateUserUseCase struct {
	userRepo domain.IUser
}

func NewUpdateUserUseCase(repo domain.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: repo}
}

func (uc *UpdateUserUseCase) Execute(id int, user entities.User) error {
	return uc.userRepo.Update(id, user)
}