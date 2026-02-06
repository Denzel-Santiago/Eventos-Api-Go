// Eventos-Api-Go/src/login/infrastructure/GetUser_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/login/application"
	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	getUserUseCase *application.GetUserUseCase
}

func NewGetUserController(getUserUseCase *application.GetUserUseCase) *GetUserController {
	return &GetUserController{getUserUseCase: getUserUseCase}
}

func (ctrl *GetUserController) Run(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := ctrl.getUserUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}