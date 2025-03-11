package routesevents

import (
	"Eventos-Api/src/eventos/infrastructure" // Importamos la infraestructura de eventos
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	// Inicializamos las dependencias de eventos
	createController, viewController, updateController, deleteController, viewAllController, getEventsByDateController := infrastructure.InitEventDependencies()

	// Grupo de rutas para eventos
	eventGroup := router.engine.Group("/events")
	{
		// Crear un evento
		eventGroup.POST("/", createController.Run)

		// Obtener un evento por ID
		eventGroup.GET("/:id", viewController.Execute)

		// Actualizar un evento por ID
		eventGroup.PUT("/:id", updateController.Execute)

		// Eliminar un evento por ID
		eventGroup.DELETE("/:id", deleteController.Run)

		// Obtener todos los eventos
		eventGroup.GET("/", viewAllController.Execute)

		// Obtener eventos por fecha
		eventGroup.GET("/date/:date", getEventsByDateController.Run)
	}
}
