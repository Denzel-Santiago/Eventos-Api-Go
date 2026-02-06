// Eventos-Api-Go/src/eventos/application/Get_UseCase.go
package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type GetEventsByDateUseCase struct {
	eventRepo domain.IEvent
}

func NewGetEventsByDateUseCase(eventRepo domain.IEvent) *GetEventsByDateUseCase {
	return &GetEventsByDateUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *GetEventsByDateUseCase) Run(date string) ([]entities.Event, error) {
	return uc.eventRepo.GetByDate(date)
}