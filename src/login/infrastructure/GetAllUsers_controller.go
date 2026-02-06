// Eventos-Api-Go/src/login/infrastructure/GetAllUsers_controller.go
package infrastructure

import (
    "net/http"

    "Eventos-Api/src/login/application"
    "github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
    getAllUsersUseCase *application.GetAllUsersUseCase
}

func NewGetAllUsersController(getAllUsersUseCase *application.GetAllUsersUseCase) *GetAllUsersController {
    return &GetAllUsersController{getAllUsersUseCase: getAllUsersUseCase}
}

func (ctrl *GetAllUsersController) Run(c *gin.Context) {
    users, err := ctrl.getAllUsersUseCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error al obtener los usuarios",
        })
        return
    }

    // ✅ CAMBIO CRÍTICO: Devolver ARRAY directo en lugar de objeto con "users"
    c.JSON(http.StatusOK, users)  // Solo el array, sin objeto wrapper
}