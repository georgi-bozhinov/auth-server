package config

import (
	"github.com/georgi-bozhinov/auth-server/server"
	"github.com/georgi-bozhinov/auth-server/server/storage"
)

type Config struct {
	DB     storage.Config
	Server server.Config
}
