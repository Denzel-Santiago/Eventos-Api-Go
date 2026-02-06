// Eventos-Api-Go/src/login/infrastructure/mysql_user_repository.go
package infrastructure

import (
	"database/sql"
	"errors"

	"Eventos-Api/src/core"
	"Eventos-Api/src/login/domain"
	"Eventos-Api/src/login/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository() domain.IUser {
	return &MySQLUserRepository{
		db: core.GetDB(),
	}
}

func (r *MySQLUserRepository) Create(user entities.User) (entities.User, error) {
	// Verificar si el usuario ya existe
	var count int
	r.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", 
		user.Username, user.Email).Scan(&count)
	
	if count > 0 {
		return entities.User{}, errors.New("el usuario o email ya existe")
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}

	query := "INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Email, string(hashedPassword), user.Role)
	if err != nil {
		return entities.User{}, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	user.Password = ""
	
	return user, nil
}

func (r *MySQLUserRepository) Update(id int, user entities.User) error {
	query := "UPDATE users SET username = ?, email = ?, role = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Role, id)
	return err
}

func (r *MySQLUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MySQLUserRepository) FindByID(id int) (entities.User, error) {
	var user entities.User
	query := "SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)
	
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entities.User{}, err
	}
	
	return user, nil
}

func (r *MySQLUserRepository) FindByUsername(username string) (entities.User, error) {
	var user entities.User
	query := "SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)
	
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entities.User{}, err
	}
	
	return user, nil
}

func (r *MySQLUserRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	query := "SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entities.User{}, err
	}
	
	return user, nil
}

func (r *MySQLUserRepository) GetAll() ([]entities.User, error) {
	query := "SELECT id, username, email, role, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func (r *MySQLUserRepository) Authenticate(username, password string) (entities.User, error) {
	// Buscar por username
	user, err := r.FindByUsername(username)
	if err != nil {
		// Si no encuentra por username, buscar por email
		user, err = r.FindByEmail(username)
		if err != nil {
			return entities.User{}, errors.New("credenciales inválidas")
		}
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entities.User{}, errors.New("credenciales inválidas")
	}

	// Limpiar contraseña antes de retornar
	user.Password = ""
	return user, nil
}