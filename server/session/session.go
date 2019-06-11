package session

import (
	"github.com/gofrs/uuid"
	"log"
	"time"
)

type Config struct {
	SessionTTL time.Duration
}

type Session struct {
	ID  string
	TTL time.Duration
}

func NewSession(cfg Config) *Session {
	id, err := uuid.NewV4()
	if err != nil {
		log.Printf("Unable to initialize session id")
		return nil
	}

	return &Session{
		ID:  id.String(),
		TTL: cfg.SessionTTL,
	}
}
