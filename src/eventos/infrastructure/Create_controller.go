// Eventos-Api-Go/src/eventos/infrastructure/Create_controller.go
package infrastructure

import (
	"net/http"

	"Eventos-Api/src/eventos/application"
	"Eventos-Api/src/eventos/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateEventController struct {
	createEventUseCase *application.CreateEventUseCase
}

func NewCreateEventController(createEventUseCase *application.CreateEventUseCase) *CreateEventController {
	return &CreateEventController{
		createEventUseCase: createEventUseCase, // CORREGIDO: minúscula
	}
}

func (ctrl *CreateEventController) Run(c *gin.Context) {
	var event entities.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del evento inválidos",
			"error":   err.Error(),
		})
		return
	}

	eventoCreado, err := ctrl.createEventUseCase.Run(event) // CORREGIDO: minúscula
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al crear el evento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Evento creado exitosamente",
		"evento":  eventoCreado,
	})
}