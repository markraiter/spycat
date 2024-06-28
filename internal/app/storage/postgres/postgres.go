package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/markraiter/spycat/internal/config"
)

type Storage struct {
	PostgresDB *sql.DB
}

func New(cfg config.Postgres) *Storage {
	// Підключення до бази даних postgres для створення нової бази даних
	initialEntryString := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.SSLMode,
	)

	initialDB, err := sql.Open(cfg.Driver, initialEntryString)
	if err != nil {
		panic(err)
	}
	defer initialDB.Close()

	_, err = initialDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.Database))
	if err != nil && err.Error() != fmt.Sprintf("pq: database \"%s\" already exists", cfg.Database) {
		panic(err)
	}

	// Підключення до новоствореної бази даних
	entryString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password,
		cfg.SSLMode,
	)

	db, err := sql.Open(cfg.Driver, entryString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return &Storage{PostgresDB: db}
}
