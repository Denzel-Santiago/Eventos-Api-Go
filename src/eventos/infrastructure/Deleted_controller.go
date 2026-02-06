// Eventos-Api-Go/src/eventos/infrastructure/Deleted_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/eventos/application"
	"github.com/gin-gonic/gin"
)

type DeleteEventController struct {
	deleteUseCase *application.DeleteEventUseCase
}

func NewDeleteEventController(deleteUseCase *application.DeleteEventUseCase) *DeleteEventController {
	return &DeleteEventController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteEventController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.deleteUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al eliminar el evento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Evento eliminado exitosamente",
	})
}