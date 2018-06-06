package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

// Database struct...
type Database struct {
	Connection *pg.DB
}

// Connect connect to database
func (db *Database) Connect() {
	db.Connection = pg.Connect(&pg.Options{
		User:     os.Getenv("DATABASE_USERNAME"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("DATABASE_PORT")),
		Database: os.Getenv("DATABASE"),
	})

	db.Connection.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()

		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})
}

// Close close a connection
func (db *Database) Close() error {
	if db.Connection != nil {
		db.Connection.Close()
		return nil
	}

	return errors.New("Can not connect database")
}

// New create a database connection.
func New() *Database {
	return &Database{}
}
