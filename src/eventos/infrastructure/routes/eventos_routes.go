// Eventos-Api-Go/src/eventos/infrastructure/routes/eventos_routes.go
package routesevents

import (
	"Eventos-Api/src/eventos/infrastructure"
	"net/http"

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
	createController, updateController, deleteController, viewAllController, getEventsByDateController := infrastructure.InitEventDependencies()

	// Grupo de rutas para eventos
	eventGroup := router.engine.Group("/events")
	{
		// ✅ Crear un evento
		eventGroup.POST("/", createController.Run)



		// ✅ Actualizar un evento por ID
		eventGroup.PUT("/:id", updateController.Execute)

		// ✅ Eliminar un evento por ID
		eventGroup.DELETE("/:id", deleteController.Run)

		// ✅ Obtener todos los eventos
		eventGroup.GET("/", viewAllController.Execute)

		// ✅ Obtener eventos por fecha
		eventGroup.GET("/date/:date", getEventsByDateController.Run)

		// ✅ Añadir manejador OPTIONS para preflight requests
		eventGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204) // Responder con No Content
		})
	}

	eventGroup.POST("/purchase", func(c *gin.Context) {
		var message map[string]interface{} // 🔹 Recibe JSON dinámico
	
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		conn, ch, err := infrastructure.ConnectRabbitMQ()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a RabbitMQ"})
			return
		}
		defer conn.Close()
		defer ch.Close()
	
		err = infrastructure.PublishTicketPurchaseMessage(ch, message) // ✅ Corrección aquí
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando mensaje a la cola"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Compra procesada y mensaje enviado a RabbitMQ correctamente"})
	})
	

	eventGroup.POST("/queue", func(c *gin.Context) {
		var message map[string]interface{} // 🔹 Cambiar tipo de datos
	
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		conn, ch, err := infrastructure.ConnectRabbitMQ()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a RabbitMQ"})
			return
		}
		defer conn.Close()
		defer ch.Close()
	
		// ✅ Enviar el mensaje completo
		err = infrastructure.PublishTicketPurchaseMessage(ch, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando mensaje a la cola"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Mensaje enviado a RabbitMQ correctamente"})
	})
	
	
}
