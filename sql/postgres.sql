-- Criar o banco de dados, caso não exista
CREATE DATABASE IF NOT EXISTS devbook;

-- Selecionar o banco de dados (não usamos 'USE' no PostgreSQL, ele é feito na conexão)
\c devbook;

-- Remover a tabela 'users' se existir
DROP TABLE IF EXISTS users;

-- Criar a tabela 'users'
CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Usamos SERIAL no PostgreSQL para auto incremento
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL UNIQUE,
    createdAt TIMESTAMP DEFAULT current_timestamp -- Sem parênteses após 'current_timestamp'
);