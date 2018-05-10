package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

type DrinkShop struct {
	Id        int64
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (shop DrinkShop) String() string {
	return fmt.Sprintf("DrinkShop<%d %s %s>", shop.Id, shop.Name, shop.Phone)
}

func main() {
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

	shop := &DrinkShop{
		Name:      "好喝飲料店",
		Phone:     "02-2730-8888",
		Address:   "台北市大安區復興南路一段321號",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := db.Insert(shop)
	if err != nil {
		panic(err)
	}
}
