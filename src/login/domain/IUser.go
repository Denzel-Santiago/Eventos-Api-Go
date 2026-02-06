// Eventos-Api-Go/src/login/domain/IUser.go
package domain

import "Eventos-Api/src/login/domain/entities"	

type IUser interface {
    Create(user entities.User) (entities.User, error)
    Update(id int, user entities.User) error
    Delete(id int) error
    FindByID(id int) (entities.User, error)
    FindByUsername(username string) (entities.User, error)
    FindByEmail(email string) (entities.User, error)
    GetAll() ([]entities.User, error)           
    GetAllBasic() ([]UserBasic, error)        
    Authenticate(username, password string) (entities.User, error)
}


type UserBasic struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}