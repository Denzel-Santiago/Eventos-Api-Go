package application

import (
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type CreateEventosUseCase struct {
	repo domain.EventoRepository
}

func NewCreateEventosUseCase(repo domain.EventoRepository) *CreateEventosUseCase {
	return &CreateEventosUseCase{repo: repo}
}

func (uc *CreateEventosUseCase) Execute(evento entities.Evento) error {
	return uc.repo.Save(evento)
}
