-- Criar o banco de dados, caso não exista
DO $$ 
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'devbook') THEN
      CREATE DATABASE devbook;
   END IF;
END $$;

-- Remover a tabela 'users' se existir
DROP TABLE IF EXISTS users;

-- Criar a tabela 'users'
CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Usamos SERIAL no PostgreSQL para auto incremento
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Usamos TIMESTAMP no PostgreSQL para data e hora
);