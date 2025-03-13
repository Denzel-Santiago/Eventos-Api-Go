package main

import (
	"fmt"

	"Eventos-Api/src/core"
	eventosRut "Eventos-Api/src/eventos/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitDB()

	r := gin.Default()

	// ✅ Registrar Middleware CORS
	r.Use(core.CORSMiddleware())

	// ✅ Configurar las rutas de eventos
	eventosRouter := eventosRut.NewRouter(r)
	eventosRouter.Run()

	fmt.Println("¡API en Funcionamiento :D!")

	// ✅ Iniciar el servidor
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
