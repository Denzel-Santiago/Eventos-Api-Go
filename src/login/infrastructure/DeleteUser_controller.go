// Eventos-Api-Go/src/login/infrastructure/DeleteUser_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/login/application"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{deleteUseCase: deleteUseCase}
}

func (ctrl *DeleteUserController) Run(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.deleteUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}