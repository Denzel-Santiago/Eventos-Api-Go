// Eventos-Api-Go/src/eventos/application/DeleteEvent_UseCase.go
package application

import (
	"Eventos-Api/src/eventos/domain"
)

type DeleteEventUseCase struct {
	eventRepo domain.IEvent
}

func NewDeleteEventUseCase(eventRepo domain.IEvent) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *DeleteEventUseCase) Run(id int) error {
	return uc.eventRepo.Delete(id)
}