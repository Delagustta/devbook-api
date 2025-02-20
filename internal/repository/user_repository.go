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

	var lastInsertedID uint64

	query := `INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	err := ur.db.QueryRow(query, user.Name, user.Nick, user.Email, user.Password).Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}
