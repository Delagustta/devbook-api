package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas na api
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Configure coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
