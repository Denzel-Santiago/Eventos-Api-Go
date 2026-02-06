// Eventos-Api-Go/src/login/infrastructure/UpdateUser_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/login/application"
	"Eventos-Api/src/login/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUseCase *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUseCase: updateUseCase}
}

func (ctrl *UpdateUserController) Run(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.updateUseCase.Execute(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}