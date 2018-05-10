package main

import (
	"drink-api/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func main() {
	router := gin.Default()

	db := pg.Connect(&pg.Options{
		User:     "kalan",
		Addr:     "192.168.99.100:3306",
		Database: "drinker-dev",
	})
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	defer db.Close()
	config.Load()
	router.Run()
}
