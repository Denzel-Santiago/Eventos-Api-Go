// Eventos-Api-Go/src/eventos/infrastructure/mysql_event_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"Eventos-Api/src/core"
	"Eventos-Api/src/eventos/domain"
	"Eventos-Api/src/eventos/domain/entities"
)

type MysqlEventRepository struct {
	conn *sql.DB
}

func NewMysqlEventRepository() domain.IEvent {
	conn := core.GetDB()
	return &MysqlEventRepository{conn: conn}
}

func (mysql *MysqlEventRepository) Save(event entities.Event) (entities.Event, error) {
	result, err := mysql.conn.Exec(
		"INSERT INTO events (name, location, date, available_tickets, price) VALUES (?, ?, ?, ?, ?)",
		event.Name,
		event.Location,
		event.Date.Format("2006-01-02"),
		event.AvailableTickets,
		event.Price,
	)
	if err != nil {
		log.Println("Error al guardar el evento:", err)
		return event, err // ✅ Retorna el evento (aunque tenga error)
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return event, err // ✅ Retorna el evento (aunque tenga error)
	}

	event.ID = int(idInserted)
	return event, nil // ✅ Ahora retorna el evento con el ID actualizado
}

func (mysql *MysqlEventRepository) Update(id int, event entities.Event) error {
	result, err := mysql.conn.Exec(
		"UPDATE events SET name = ?, location = ?, date = ?, available_tickets = ?, price = ? WHERE id = ?",
		event.Name,
		event.Location,
		event.Date.Format("2006-01-02"),
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
	var dateStr string
	
	row := mysql.conn.QueryRow(
		"SELECT id, name, location, date, available_tickets, price, created_at FROM events WHERE id = ?", 
		id,
	)

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&dateStr,
		&event.AvailableTickets,
		&event.Price,
		&event.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Event{}, fmt.Errorf("evento con ID %d no encontrado", id)
		}
		return entities.Event{}, err
	}

	// Parsear la fecha
	event.Date, _ = time.Parse("2006-01-02", dateStr)
	return event, nil
}

func (mysql *MysqlEventRepository) GetAll() ([]entities.Event, error) {
	var events []entities.Event

	rows, err := mysql.conn.Query(
		"SELECT id, name, location, date, available_tickets, price, created_at FROM events",
	)
	
	if err != nil {
		log.Println("Error al obtener todos los eventos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		var dateStr string
		
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&dateStr,
			&event.AvailableTickets,
			&event.Price,
			&event.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear evento:", err)
			return nil, err
		}
		
		// Parsear la fecha
		event.Date, _ = time.Parse("2006-01-02", dateStr)
		events = append(events, event)
	}

	return events, nil
}

func (mysql *MysqlEventRepository) GetByDate(date string) ([]entities.Event, error) {
	var events []entities.Event
	
	rows, err := mysql.conn.Query(
		"SELECT id, name, location, date, available_tickets, price, created_at FROM events WHERE date = ?", 
		date,
	)
	
	if err != nil {
		log.Println("Error al obtener eventos por fecha:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		var dateStr string
		
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&dateStr,
			&event.AvailableTickets,
			&event.Price,
			&event.CreatedAt,
		)
		if err != nil {
			log.Println("Error al filtrar los eventos:", err)
			return nil, err
		}
		
		// Parsear la fecha
		event.Date, _ = time.Parse("2006-01-02", dateStr)
		events = append(events, event)
	}

	return events, nil
}

func (mysql *MysqlEventRepository) FindByLocation(location string) ([]entities.Event, error) {
	var events []entities.Event
	
	rows, err := mysql.conn.Query(
		"SELECT id, name, location, date, available_tickets, price, created_at FROM events WHERE location LIKE ?", 
		"%"+location+"%",
	)
	
	if err != nil {
		log.Println("Error al obtener eventos por ubicación:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		var dateStr string
		
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Location,
			&dateStr,
			&event.AvailableTickets,
			&event.Price,
			&event.CreatedAt,
		)
		if err != nil {
			log.Println("Error al filtrar eventos por ubicación:", err)
			return nil, err
		}
		
		// Parsear la fecha
		event.Date, _ = time.Parse("2006-01-02", dateStr)
		events = append(events, event)
	}

	return events, nil
}