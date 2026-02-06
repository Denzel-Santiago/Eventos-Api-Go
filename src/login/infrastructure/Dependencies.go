// Eventos-Api-Go/src/login/infrastructure/Dependencies.go
package infrastructure

import (
	"Eventos-Api/src/login/application"
)

func InitUserDependencies() (
	*LoginController,
	*CreateUserController,
	*GetUserController,
	*UpdateUserController,
	*DeleteUserController,
) {
	// Repositorio
	userRepo := NewMySQLUserRepository()
	
	// Use Cases
	loginUseCase := application.NewLoginUseCase(userRepo)
	createUseCase := application.NewCreateUserUseCase(userRepo)
	getUserUseCase := application.NewGetUserUseCase(userRepo)
	updateUseCase := application.NewUpdateUserUseCase(userRepo)
	deleteUseCase := application.NewDeleteUserUseCase(userRepo)
	
	// Controladores
	loginController := NewLoginController(loginUseCase)
	createController := NewCreateUserController(createUseCase)
	getUserController := NewGetUserController(getUserUseCase)
	updateController := NewUpdateUserController(updateUseCase)
	deleteController := NewDeleteUserController(deleteUseCase)
	
	return loginController, createController, getUserController, updateController, deleteController
}