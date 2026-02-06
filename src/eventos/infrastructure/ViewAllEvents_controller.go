// Eventos-Api-Go/src/eventos/infrastructure/ViewAllEvents_controller.go
package infrastructure

import (
	"net/http"

	"Eventos-Api/src/eventos/application"
	"github.com/gin-gonic/gin"
)

type ViewAllEventsController struct {
	useCase *application.ViewAllEvents
}

func NewViewAllEventsController(useCase *application.ViewAllEvents) *ViewAllEventsController {
	return &ViewAllEventsController{useCase: useCase}
}

func (ctrl *ViewAllEventsController) Execute(c *gin.Context) {
	events, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener los eventos",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}