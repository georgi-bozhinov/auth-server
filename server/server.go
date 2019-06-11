package server

import (
	"github.com/georgi-bozhinov/auth-server/server/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Config struct {
	Port string
}

type Server struct {
	Router *mux.Router
	Config Config
}

func NewServer(cfg Config, api *api.Api) *Server {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	return &Server{Router: router, Config: cfg}
}

func (s *Server) Start() error {
	log.Printf("Auth server starting on port %s", s.Config.Port)
	return http.ListenAndServe(":"+s.Config.Port, s.Router)
}
