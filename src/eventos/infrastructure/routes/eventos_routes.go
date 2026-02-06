// Eventos-Api-Go/src/eventos/infrastructure/routes/eventos_routes.go
package routes

import (
	"Eventos-Api/src/eventos/infrastructure"
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
	createController, updateController, deleteController, viewAllController, viewEventController, getEventsByDateController := infrastructure.InitEventDependencies()

	eventGroup := router.engine.Group("/events")
	{
		// CRUD de eventos
		eventGroup.POST("/", createController.Run)
		eventGroup.GET("/", viewAllController.Execute)
		eventGroup.GET("/:id", viewEventController.Run)
		eventGroup.PUT("/:id", updateController.Execute)
		eventGroup.DELETE("/:id", deleteController.Run)
		
		// Búsquedas específicas
		eventGroup.GET("/date/:date", getEventsByDateController.Run)
		
		// Para CORS
		eventGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204)
		})
	}
}