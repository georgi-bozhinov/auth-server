package api

import (
	"github.com/georgi-bozhinov/auth-server/server/api/users"
	"github.com/georgi-bozhinov/auth-server/server/router"
	"github.com/georgi-bozhinov/auth-server/server/storage"
	"github.com/gorilla/mux"
)

type Api struct {
	Routes []router.Route
}

func NewAPI(storage storage.Storage) *Api {
	userController := users.Controller{Storage: storage}

	return &Api{Routes: userController.Routes()}
}

func (a *Api) RegisterRoutes(router *mux.Router) {
	for _, route := range a.Routes {
		router.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(route.Handler)
	}
}
