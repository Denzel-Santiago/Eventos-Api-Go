package infrastructure

import (
	"Eventos-Api/src/eventos/application"
	"log"
)

func InitEventDependencies() (
	*CreateEventController,
	*ViewEventController,
	*UpdateEventController,
	*DeleteEventController,
	*ViewAllEventsController,
	*GetEventsByDateController,
) {

	repo := NewMysqlEventRepository()

	createUseCase := application.NewCreateEventUseCase(repo)
	viewUseCase := application.NewViewEvent(repo)
	updateUseCase := application.NewUpdateEvent(repo)
	deleteUseCase := application.NewDeleteEventUseCase(repo)
	viewAllUseCase := application.NewViewAllEvents(repo)
	getEventsByDate := application.NewGetEventsByDateUseCase(repo)

	// ✅ Conexión a RabbitMQ
	conn, ch, err := ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	// ✅ Declarar la cola (asegúrate de que el nombre sea correcto)
	queueName := "mi-cola"
	_, err = DeclareQueue(ch, queueName)
	if err != nil {
		log.Fatalf("❌ Error al declarar la cola: %v", err)
	}

	// Crear controladores
	createController := NewCreateEventController(createUseCase)
	viewController := NewViewEventController(viewUseCase)
	updateController := NewUpdateEventController(updateUseCase)
	deleteController := NewDeleteEventController(deleteUseCase)
	viewAllController := NewViewAllEventsController(viewAllUseCase)
	getEventsByDateController := NewGetEventsByDateController(getEventsByDate)

	return createController, viewController, updateController, deleteController, viewAllController, getEventsByDateController
}
