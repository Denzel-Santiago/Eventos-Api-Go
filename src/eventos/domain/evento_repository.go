package domain

import (
	"Eventos-Api/src/eventos/domain/entities"
	"database/sql"
)

type EventoRepository interface {
	Save(evento entities.Evento) error
}

type eventoRepository struct {
	db *sql.DB
}

func NewEventoRepository(db *sql.DB) EventoRepository {
	return &eventoRepository{db: db}
}

func (r *eventoRepository) Save(evento entities.Evento) error {
	_, err := r.db.Exec("INSERT INTO events (name, location, date, available_tickets, price) VALUES (?, ?, ?, ?, ?)",
		evento.Name, evento.Location, evento.Date, evento.AvailableTickets, evento.Price)
	return err
}
