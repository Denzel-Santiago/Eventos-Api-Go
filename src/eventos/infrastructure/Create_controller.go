package infrastructure

import (
	"Eventos-Api/src/eventos/application"
	"Eventos-Api/src/eventos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EventosController maneja las solicitudes HTTP para los eventos.
type EventosController struct {
	createEventoUseCase *application.CreateEventosUseCase
}

// NewEventosController crea una nueva instancia de EventosController.
func NewEventosController(createEventoUseCase *application.CreateEventosUseCase) *EventosController {
	return &EventosController{createEventoUseCase: createEventoUseCase}
}

// CreateEvento maneja la solicitud para crear un nuevo evento.
func (c *EventosController) CreateEvento(ctx *gin.Context) {
	var evento entities.Evento
	if err := ctx.ShouldBindJSON(&evento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.createEventoUseCase.Execute(evento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Evento created successfully"})
}
