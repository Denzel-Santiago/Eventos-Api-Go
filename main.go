// Eventos-Api-Go/main.go
package main

import (
	"fmt"
	"log"

	"Eventos-Api/src/core"
	eventosRut "Eventos-Api/src/eventos/infrastructure/routes"
	loginRut "Eventos-Api/src/login/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar base de datos
	core.InitDB()

	// Crear usuario admin por defecto si no existe
	err := createDefaultAdmin()
	if err != nil {
		log.Printf("Advertencia al crear admin: %v", err)
	}

	// Configurar Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(core.CORSMiddleware())

	// Configurar rutas de eventos
	eventosRouter := eventosRut.NewRouter(r)
	eventosRouter.Run()

	// Configurar rutas de login
	loginRouter := loginRut.NewLoginRouter(r)
	loginRouter.Run()

	fmt.Println("¡API en Funcionamiento en puerto 8000! 🚀")

	// Iniciar servidor
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}

func createDefaultAdmin() error {
	db := core.GetDB()
	
	// Verificar si ya existe el usuario admin
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Crear usuario admin con password: admin123
		hashedPassword := "$2a$10$N9qo8uLOickgx2ZMRZoMye.CmJ3wL3kH9b6W7Bq6K8p7Qc2Yz8KZa" // admin123 hasheado
		_, err := db.Exec(
			"INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)",
			"admin", "admin@eventos.com", hashedPassword, "admin",
		)
		return err
	}

	return nil
}