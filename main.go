package main

import (
	"github.com/georgi-bozhinov/auth-server/server"
	"github.com/georgi-bozhinov/auth-server/server/api"
	"github.com/georgi-bozhinov/auth-server/server/config"
	"github.com/georgi-bozhinov/auth-server/server/storage"
	"log"
)

func BuildAuthServer(db storage.Storage) *server.Server {
	a := api.NewAPI(db)
	s := server.NewServer(server.Config{Port: "8000"}, a)

	return s
}

func main() {
	// config := readConfigFromFile(file)
	// config.DB
	// config.Server

	cfg := config.Config{DB: storage.Config{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		DbName:   "authserver",
		Type:     "postgres",
	}}

	db, err := storage.New(cfg.DB)

	if err != nil {
		log.Fatalf("Unable to initialize db: %v", err)
	}

	srv := BuildAuthServer(db)

	if err := srv.Start(); err != nil {
		log.Fatalf("Unable to initialize auth server: %v", err)
	}
}
