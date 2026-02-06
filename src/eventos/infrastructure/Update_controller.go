// Eventos-Api-Go/src/eventos/infrastructure/Update_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/eventos/application"
	"Eventos-Api/src/eventos/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateEventController struct {
	useCase *application.UpdateEvent
}

func NewUpdateEventController(useCase *application.UpdateEvent) *UpdateEventController {
	return &UpdateEventController{useCase: useCase}
}

func (ctrl *UpdateEventController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var event entities.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	err = ctrl.useCase.Execute(id, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al actualizar el evento",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Evento actualizado exitosamente",
	})
}