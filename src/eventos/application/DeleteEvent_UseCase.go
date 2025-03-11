package application

import "Eventos-Api/src/eventos/domain"

type DeleteEventUseCase struct {
	db domain.IEvent
}

func NewDeleteEventUseCase(db domain.IEvent) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		db: db,
	}
}

func (uc *DeleteEventUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
