package main

import (
	"Eventos-Api/src/eventos/infrastructure"
	"Eventos-Api/src/eventos/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configura las dependencias
	deps, err := infrastructure.SetupDependencies()
	if err != nil {
		log.Fatalf("Failed to setup dependencies: %v", err)
	}

	// Conectar a RabbitMQ y enviar un mensaje de prueba
	conn, ch, err := infrastructure.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Verificar que la cola 'queue' existe
	_, err = infrastructure.DeclareQueue(ch, "queue")
	if err != nil {
		log.Fatalf("Error al declarar la cola: %v", err)
	}

	// Enviar mensaje de prueba
	err = infrastructure.PublishTestMessage(ch, "queue")
	if err != nil {
		log.Fatalf("Error al enviar mensaje de prueba: %v", err)
	}

	// Inicializa el router de Gin
	r := gin.Default()

	// Configura las rutas pasando el handler de eventos
	routes.SetupEventosRoutes(r, deps.EventosHandler)

	// Inicia el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
