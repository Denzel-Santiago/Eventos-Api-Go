package infrastructure

import (
	"net/http"
	"strconv"

	application "Eventos-Api/src/eventos/application"
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

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el evento",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Evento eliminado exitosamente",
	})
}
