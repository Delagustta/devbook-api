package repository

import (
	"api/internal/model"
	"database/sql"
	"fmt"
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

// FindBy busca todos os usuários que atendem o filtro de nome ou nick
func (ur userRepository) FindBy(nameOrNick string) ([]model.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	query := `SELECT id, name, nick, email, created_at FROM users WHERE name ILIKE $1 OR nick ILIKE $1`

	rows, err := ur.db.Query(query, nameOrNick)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID busca um usuário pelo ID
func (ur userRepository) FindByID(id uint64) (model.User, error) {

	query := `SELECT id, name, nick, email, created_at FROM users WHERE id = $1`

	rows, err := ur.db.Query(query, id)
	if err != nil {
		return model.User{}, err
	}

	defer rows.Close()

	var user model.User

	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
