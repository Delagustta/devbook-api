package controller

import "net/http"

// CreateUser insere um novo usuario no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

// GetAllUsers busca todos os usuarios
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

// FindUser busca um usuario
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

// UpdateUser atualiza um usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

// DeleteUser deleta um usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
