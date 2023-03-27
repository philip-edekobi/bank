package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/philip-edekobi/bank/util"

	db "github.com/philip-edekobi/bank/db/sqlc"

	"github.com/philip-edekobi/bank/api"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database", err)
	}

	store := db.NewStore(conn)
	server, serverErr := api.NewServer(config, store)
	if serverErr != nil {
		log.Fatal(fmt.Errorf("failed to start server: %w", serverErr))
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
