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
func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DATABASE_USERNAME"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("DATABASE_PORT")),
		Database: os.Getenv("DATABASE"),
	})

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()

		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return db
}

// Close close a connection
func Close(db *pg.DB) error {
	db.Close()

	return errors.New("Can not connect database")
}
