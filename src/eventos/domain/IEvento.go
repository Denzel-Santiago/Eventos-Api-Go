package domain

import (
	"Eventos-Api/src/eventos/domain/entities"
)

type IEvent interface {
	Save(event entities.Event) error
	Update(id int, event entities.Event) error
	Delete(id int) error
	FindByID(id int) (entities.Event, error)
	GetAll() ([]entities.Event, error)
	GetByDate(date string) ([]entities.Event, error)
}
