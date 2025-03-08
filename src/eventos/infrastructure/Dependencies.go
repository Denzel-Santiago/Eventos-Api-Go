package infrastructure

import (
	"Eventos-Api/src/core"
	"Eventos-Api/src/eventos/application"
	"Eventos-Api/src/eventos/domain"
)

// Dependencies contiene todas las dependencias necesarias para la aplicación.
type Dependencies struct {
	DB                  *core.DB
	EventoRepository    domain.EventoRepository
	CreateEventoUseCase *application.CreateEventosUseCase
	EventosHandler      *application.EventosHandler
}

// SetupDependencies configura y devuelve las dependencias de la aplicación.
func SetupDependencies() (*Dependencies, error) {
	// Configura la conexión a la base de datos
	db, err := core.NewDB()
	if err != nil {
		return nil, err
	}

	// Inicializa el repositorio de eventos
	eventoRepo := domain.NewEventoRepository(db.DB) // Pasamos db.DB (el *sql.DB subyacente)

	// Inicializa el caso de uso para crear eventos
	createEventoUseCase := application.NewCreateEventosUseCase(eventoRepo)

	// Inicializa el handler de eventos
	eventosHandler := application.NewEventosHandler(createEventoUseCase)

	// Retorna las dependencias configuradas
	return &Dependencies{
		DB:                  db,
		EventoRepository:    eventoRepo,
		CreateEventoUseCase: createEventoUseCase,
		EventosHandler:      eventosHandler,
	}, nil
}
