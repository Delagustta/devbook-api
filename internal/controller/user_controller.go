package controller

import "net/http"

// CriarUsuario insere um novo usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}

// BuscarTodosUsuarios busca todos os usuarios
func BuscarTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

// BuscarUsuario busca um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CBuscando um Usuario"))
}

// AtualizarUsuario atualiza um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario"))
}

// DeletarUsuario deleta um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario"))
}
