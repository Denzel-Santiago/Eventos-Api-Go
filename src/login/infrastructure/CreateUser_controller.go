// Eventos-Api-Go/src/login/infrastructure/CreateUser_controller.go
package infrastructure

import (
	"net/http"

	"Eventos-Api/src/login/application"
	"Eventos-Api/src/login/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUseCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{createUseCase: createUseCase}
}

func (ctrl *CreateUserController) Run(c *gin.Context) {
	var user entities.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer rol por defecto si no se especifica
	if user.Role == "" {
		user.Role = "user"
	}

	createdUser, err := ctrl.createUseCase.Execute(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}