package main

import (
	"log"

	"github.com/ppp3ppj/pppfav-htmx/config"
	"github.com/ppp3ppj/pppfav-htmx/db"
	"github.com/ppp3ppj/pppfav-htmx/internal/server"
)

func main() {
    conf := config.ConfigGetting()
    db := db.NewPostgresDatabase(conf.Database)
    defer func() {
        if err := db.Close(); err != nil {
            log.Fatal("Failed to close database connection: %v", err)
        }
    }()

    server := server.NewEchoServer(conf, db)
    server.Start()
}
