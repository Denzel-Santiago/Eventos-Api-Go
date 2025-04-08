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

	createController, updateController, deleteController, viewAllController, getEventsByDateController := infrastructure.InitEventDependencies()

	eventGroup := router.engine.Group("/events")
	{

		eventGroup.POST("/", createController.Run)

		eventGroup.PUT("/:id", updateController.Execute)

		eventGroup.DELETE("/:id", deleteController.Run)

		eventGroup.GET("/", viewAllController.Execute)

		eventGroup.GET("/date/:date", getEventsByDateController.Run)

		eventGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204)
		})
	}

	eventGroup.POST("/purchase", func(c *gin.Context) {
		var message map[string]interface{}

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

		err = infrastructure.PublishTicketPurchaseMessage(ch, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando mensaje a la cola"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Compra procesada y mensaje enviado a RabbitMQ correctamente"})
	})

	eventGroup.POST("/queue", func(c *gin.Context) {
		var message map[string]interface{}

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

		err = infrastructure.PublishTicketPurchaseMessage(ch, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando mensaje a la cola"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Mensaje enviado a RabbitMQ correctamente"})
	})

}
