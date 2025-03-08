package routes

import (
	"Eventos-Api/src/eventos/application"
	"github.com/gin-gonic/gin"
)

func SetupEventosRoutes(r *gin.Engine, eventosHandler *application.EventosHandler) {
	r.POST("/eventos", eventosHandler.CreateEvento)
	r.GET("/eventos", eventosHandler.GetEventos)
	// Otras rutas...
}
