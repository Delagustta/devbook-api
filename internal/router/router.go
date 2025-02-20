package router

import (
	"api/internal/router/route"

	"github.com/gorilla/mux"
)

// Generate retorna um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return route.Configure(r)
}
