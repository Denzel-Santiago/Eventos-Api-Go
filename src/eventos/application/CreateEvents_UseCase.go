package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type CreateEventUseCase struct {
	db domain.IEvent
}

func NewCreateEventUseCase(db domain.IEvent) *CreateEventUseCase {
	return &CreateEventUseCase{
		db: db,
	}
}

func (uc *CreateEventUseCase) Run(event *entities.Event) (*entities.Event, error) {
	err := uc.db.Save(*event)
	return event, err
}
