package database

import (
	"api/internal/config"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // Driver para conexão com o banco de dados
)

// Connect realiza a conexão com o banco de dados
func Connect() (*sql.DB, error) {

	db, err := sql.Open("pgx", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
