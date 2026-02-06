// Eventos-Api-Go/src/login/routes/login_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"Eventos-Api/src/login/infrastructure"
)

type LoginRouter struct {
	engine *gin.Engine
}

func NewLoginRouter(engine *gin.Engine) *LoginRouter {
	return &LoginRouter{
		engine: engine,
	}
}


func (router *LoginRouter) Run() {
	loginController, createController, getUserController, updateController, deleteController := infrastructure.InitUserDependencies()

	authGroup := router.engine.Group("/auth")
	{
		// Autenticación
		authGroup.POST("/login", loginController.Run)
		
		// Gestión de usuarios
		authGroup.POST("/register", createController.Run)
		authGroup.GET("/users/:id", getUserController.Run)
		authGroup.PUT("/users/:id", updateController.Run)
		authGroup.DELETE("/users/:id", deleteController.Run)
		
		// Para CORS
		authGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204)
		})
	}
}