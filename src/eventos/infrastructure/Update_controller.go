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

    var updateData map[string]interface{}
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // Obtener evento actual
    currentEvent, err := uuc.useCase.Db.FindByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
        return
    }

    // Verificar si es una operación de decremento
    if operation, ok := updateData["operation"].(string); ok && operation == "decrement" {
        if currentEvent.AvailableTickets <= 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No hay boletos disponibles"})
            return
        }
        currentEvent.AvailableTickets -= 1
    } else {
        // Actualización normal de campos
        if name, ok := updateData["name"].(string); ok {
            currentEvent.Name = name
        }
        if location, ok := updateData["location"].(string); ok {
            currentEvent.Location = location
        }
        // ... otros campos si son necesarios
    }

    // Actualizar evento en la BD
    err = uuc.useCase.Execute(id, currentEvent)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el evento"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Evento actualizado exitosamente",
        "data":    currentEvent,
    })
}