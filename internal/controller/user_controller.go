package controller

import (
	"api/internal/database"
	"api/internal/model"
	"api/internal/repository"
	"api/internal/response"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CreateUser insere um novo usuario no banco
func CreateUser(httpWriter http.ResponseWriter, httpRequest *http.Request) {

	requestBody, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		response.Error(httpWriter, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(httpWriter, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		response.Error(httpWriter, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(httpWriter, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, err = repository.CreateUser(user)
	if err != nil {
		response.Error(httpWriter, http.StatusInternalServerError, err)
		return
	}

	response.JSON(httpWriter, http.StatusCreated, user)
}

// FindUsersBy busca todos os usuarios que atendem o filtro de nome ou nick
func FindUsersBy(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, err := repository.FindBy(nameOrNick)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
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
