// Eventos-Api-Go/src/eventos/infrastructure/Get_controller.go
package infrastructure

import (
	"net/http"

	"Eventos-Api/src/eventos/application"
	"github.com/gin-gonic/gin"
)

type GetEventsByDateController struct {
	getByDateUseCase *application.GetEventsByDateUseCase
}

func NewGetEventsByDateController(getByDateUseCase *application.GetEventsByDateUseCase) *GetEventsByDateController {
	return &GetEventsByDateController{
		getByDateUseCase: getByDateUseCase,
	}
}

func (ctrl *GetEventsByDateController) Run(c *gin.Context) {
	date := c.Param("date")

	events, err := ctrl.getByDateUseCase.Run(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener eventos por fecha",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}