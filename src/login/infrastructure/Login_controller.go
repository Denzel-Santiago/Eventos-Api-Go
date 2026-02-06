// Eventos-Api-Go/src/login/infrastructure/Login_controller.go
package infrastructure

import (
	"net/http"

	"Eventos-Api/src/login/application"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUseCase *application.LoginUseCase
}

func NewLoginController(loginUseCase *application.LoginUseCase) *LoginController {
	return &LoginController{loginUseCase: loginUseCase}
}

func (ctrl *LoginController) Run(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.loginUseCase.Execute(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    user,
	})
}