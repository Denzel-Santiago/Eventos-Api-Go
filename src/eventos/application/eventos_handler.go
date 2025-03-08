package application

import (
	"Eventos-Api/src/eventos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventosHandler struct {
	createEventoUseCase *CreateEventosUseCase
}

func NewEventosHandler(createEventoUseCase *CreateEventosUseCase) *EventosHandler {
	return &EventosHandler{createEventoUseCase: createEventoUseCase}
}

func (h *EventosHandler) CreateEvento(ctx *gin.Context) {
	var evento entities.Evento
	if err := ctx.ShouldBindJSON(&evento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createEventoUseCase.Execute(evento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Evento created successfully"})
}

func (h *EventosHandler) GetEventos(ctx *gin.Context) {
	// Lógica para obtener eventos
	ctx.JSON(http.StatusOK, gin.H{"message": "List of eventos"})
}
