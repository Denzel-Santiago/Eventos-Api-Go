// Eventos-Api-Go/src/eventos/infrastructure/Dependencies.go
package infrastructure

import (
	"Eventos-Api/src/eventos/application"
)

func InitEventDependencies() (
	*CreateEventController,
	*UpdateEventController,
	*DeleteEventController,
	*ViewAllEventsController,
	*ViewEventController,
	*GetEventsByDateController,
) {
	// Repositorio
	repo := NewMysqlEventRepository()

	// Use Cases
	createUseCase := application.NewCreateEventUseCase(repo)
	updateUseCase := application.NewUpdateEvent(repo)
	deleteUseCase := application.NewDeleteEventUseCase(repo)
	viewAllUseCase := application.NewViewAllEvents(repo)
	viewEventUseCase := application.NewViewEvent(repo)
	getEventsByDate := application.NewGetEventsByDateUseCase(repo)

	// Controladores
	createController := NewCreateEventController(createUseCase)
	updateController := NewUpdateEventController(updateUseCase)
	deleteController := NewDeleteEventController(deleteUseCase)
	viewAllController := NewViewAllEventsController(viewAllUseCase)
	viewEventController := NewViewEventController(viewEventUseCase)
	getEventsByDateController := NewGetEventsByDateController(getEventsByDate)

	return createController, updateController, deleteController, viewAllController, viewEventController, getEventsByDateController
}