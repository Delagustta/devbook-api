package route

import (
	"api/internal/controller"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Function:     controller.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/user",
		Method:       http.MethodGet,
		Function:     controller.GetAllUsers,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodGet,
		Function:     controller.FindUser,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodPut,
		Function:     controller.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodDelete,
		Function:     controller.DeleteUser,
		AuthRequired: false,
	},
}
