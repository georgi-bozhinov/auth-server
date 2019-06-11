package storage

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// Storage ...
type Storage interface {
	DB() *sqlx.DB
	Close() error
}

type Config struct {
	Type     string
	Host     string
	User     string
	Password string
	DbName   string
}

type postgresStorage struct {
	db *sqlx.DB
}

func New(cfg Config) (Storage, error) {
	connString := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", cfg.Type, cfg.User, cfg.Password, cfg.Host, cfg.DbName)

	db, err := sqlx.Open(cfg.Type, connString)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://server/migrations", cfg.DbName, driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Database up to date.")
			err = nil
		} else {
			return nil, err
		}
	}

	return &postgresStorage{db: db}, nil
}

func (ps *postgresStorage) DB() *sqlx.DB {
	return ps.db
}

func (ps *postgresStorage) Close() error {
	if err := ps.db.Close(); err != nil {
		return err
	}

	return nil
}
