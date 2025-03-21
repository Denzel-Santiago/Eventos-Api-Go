// Eventos-Api-Go/src/eventos/infrastructure/mysql_event_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"Eventos-Api/src/core"
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type MysqlEventRepository struct {
	conn *sql.DB
}

// FindByLocation implements domain.IEvent.
func (mysql *MysqlEventRepository) FindByLocation(location string) ([]entities.Event, error) {
	panic("unimplemented")
}

func NewMysqlEventRepository() domain.IEvent {
	conn := core.GetDB()
	return &MysqlEventRepository{conn: conn}
}

func (mysql *MysqlEventRepository) Save(event entities.Event) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO events (name, location, date, available_tickets, price) VALUES (?, ?, ?, ?, ?)",
		event.Name,
		event.Location,
		event.Date,
		event.AvailableTickets,
		event.Price,
	)
	if err != nil {
		log.Println("Error al guardar el evento:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	event.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlEventRepository) Update(id int, event entities.Event) error {
	result, err := mysql.conn.Exec(
		"UPDATE events SET name = ?, location = ?, date = ?, available_tickets = ?, price = ? WHERE id = ?",
		event.Name,
		event.Location,
		event.Date,
		event.AvailableTickets,
		event.Price,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el evento:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el evento con ID:", id)
		return fmt.Errorf("evento con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlEventRepository) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el evento:", err)
		return err
	}
	return nil
}

func (mysql *MysqlEventRepository) FindByID(id int) (entities.Event, error) {
	var event entities.Event
	row := mysql.conn.QueryRow("SELECT id, name, location, date, available_tickets, price, created_at FROM events WHERE id = ?", id)

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Date,
		&event.AvailableTickets,
		&event.Price,
		&event.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Evento no encontrado:", err)
			return entities.Event{}, fmt.Errorf("evento con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el evento por ID:", err)
		return entities.Event{}, err
	}

	return event, nil
}

func (mysql *MysqlEventRepository) GetAll() ([]entities.Event, error) {
	var events []entities.Event

	rows, err := mysql.conn.Query("SELECT id, name, location, date, available_tickets, price, CreatedAt FROM events")
	if err != nil {
		log.Println("Error al obtener todos los eventos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&event.Date,
			&event.AvailableTickets,
			&event.Price,
			&event.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear evento:", err)
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return events, nil
}

func (mysql *MysqlEventRepository) GetByDate(date string) ([]entities.Event, error) {
	var events []entities.Event

	rows, err := mysql.conn.Query("SELECT id, name, location, date, available_tickets, price, created_at FROM events WHERE date = ?", date)
	if err != nil {
		log.Println("Error al obtener eventos por fecha:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&event.Date,
			&event.AvailableTickets,
			&event.Price,
			&event.CreatedAt,
		)
		if err != nil {
			log.Println("Error al filtrar los eventos:", err)
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al filtrar los eventos:", err)
		return nil, err
	}

	return events, nil
}
