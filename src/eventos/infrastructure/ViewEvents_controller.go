package infrastructure

import (
	"net/http"
	"strconv"

	"Eventos-Api/src/eventos/application"
	"github.com/gin-gonic/gin"
)

type ViewEventController struct {
	useCase *application.ViewEvent
}

func NewViewEventController(useCase *application.ViewEvent) *ViewEventController {
	return &ViewEventController{useCase: useCase}
}

func (vec *ViewEventController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	event, err := vec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	c.JSON(http.StatusOK, event)
}
