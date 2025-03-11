package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type GetEventsByDateUseCase struct {
	db domain.IEvent
}

func NewGetEventsByDateUseCase(db domain.IEvent) *GetEventsByDateUseCase {
	return &GetEventsByDateUseCase{
		db: db,
	}
}

func (uc *GetEventsByDateUseCase) Run(date string) ([]entities.Event, error) {
	return uc.db.GetByDate(date)
}
