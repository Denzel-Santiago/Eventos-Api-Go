package infrastructure

import (
	"net/http"

	"Eventos-Api/src/eventos/application"
	"Eventos-Api/src/eventos/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateEventController struct {
	CreateEventUseCase *application.CreateEventUseCase
}

func NewCreateEventController(createEventUseCase *application.CreateEventUseCase) *CreateEventController {
	return &CreateEventController{
		CreateEventUseCase: createEventUseCase,
	}
}

func (ctrl *CreateEventController) Run(c *gin.Context) {
	var event entities.Event

	if errJSON := c.ShouldBindJSON(&event); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del evento inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	eventoCreado, errAdd := ctrl.CreateEventUseCase.Run(&event)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el evento",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "El evento ha sido agregado",
		"evento":  eventoCreado,
	})
}
