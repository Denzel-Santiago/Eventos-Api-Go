// Eventos-Api-Go/src/login/application/GetAllUsers_UseCase.go
package application

import (
    "Eventos-Api/src/login/domain"
)

type GetAllUsersUseCase struct {
    userRepo domain.IUser
}

func NewGetAllUsersUseCase(repo domain.IUser) *GetAllUsersUseCase {
    return &GetAllUsersUseCase{userRepo: repo}
}

func (uc *GetAllUsersUseCase) Execute() ([]domain.UserBasic, error) {
    return uc.userRepo.GetAllBasic()  // Esto ya devuelve []domain.UserBasic
}