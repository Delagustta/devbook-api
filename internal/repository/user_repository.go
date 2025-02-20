package repository

import (
	"api/internal/model"
	"database/sql"
)

// userRepository é uma estrutura para manipulação de dados de usuários
type userRepository struct {
	db *sql.DB
}

// newUserRepository cria um novo userRepository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

// CreateUser insere um novo usuário no banco
func (ur userRepository) CreateUser(user model.User) (uint64, error) {

	statement, err := ur.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
