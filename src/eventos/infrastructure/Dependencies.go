package infrastructure

import (
	"Eventos-Api/src/eventos/application"
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

	// Crear controladores
	createController := NewCreateEventController(createUseCase)
	viewController := NewViewEventController(viewUseCase)
	updateController := NewUpdateEventController(updateUseCase)
	deleteController := NewDeleteEventController(deleteUseCase)
	viewAllController := NewViewAllEventsController(viewAllUseCase)
	getEventsByDateController := NewGetEventsByDateController(getEventsByDate)

	return createController, viewController, updateController, deleteController, viewAllController, getEventsByDateController
}
