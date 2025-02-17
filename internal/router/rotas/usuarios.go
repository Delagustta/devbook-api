package rotas

import (
	"api/internal/controller"
	"net/http"
)

var userRoutes = []Rota{
	{
		URI:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controller.CriarUsuario,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controller.BuscarTodosUsuarios,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodGet,
		Funcao:       controller.BuscarUsuario,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodPut,
		Funcao:       controller.AtualizarUsuario,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodDelete,
		Funcao:       controller.DeletarUsuario,
		AuthRequired: false,
	},
}
