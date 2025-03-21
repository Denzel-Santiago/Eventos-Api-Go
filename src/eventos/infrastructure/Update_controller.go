//Eventos-Api-Go/src/eventos/infrastructure/Update_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/eventos/application"

	"github.com/gin-gonic/gin"
)

type UpdateEventController struct {
	useCase *application.UpdateEvent
}

func NewUpdateEventController(useCase *application.UpdateEvent) *UpdateEventController {
	return &UpdateEventController{useCase: useCase}
}

func (uuc *UpdateEventController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	var updateData struct {
		AvailableTickets int `json:"available_tickets"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// 📌 Obtener evento actual para verificar si existe
	event, err := uuc.useCase.Db.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	// 📌 Asegurar que haya boletos disponibles antes de restar
	if event.AvailableTickets+updateData.AvailableTickets < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay suficientes boletos disponibles"})
		return
	}

	// 📌 Aplicar la reducción de boletos
	event.AvailableTickets += updateData.AvailableTickets

	// 📌 Actualizar evento en la BD
	err = uuc.useCase.Execute(id, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Boleto actualizado exitosamente"})
}
