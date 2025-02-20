package controller

import (
	"api/internal/database"
	"api/internal/model"
	"api/internal/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser insere um novo usuario no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {

	request, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err = json.Unmarshal(request, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUserRepository(db)
	id, err := repository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso. ID: %d", id)))
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
